package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang-demo/video-go/video-api/config"
	"golang.org/x/sys/windows"
	"math"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

// GetMD5 获取目标字符串加盐的md5加密字符串
func GetMD5(targetStr string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(targetStr + config.VideoConf.MD5Salt))
	ctxStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(ctxStr)
}

// GetUUid 获取全局唯一uuid
// 默认规则 机器码 + 时间戳 + 用户唯一标示 + 随机数
func GetUUid(userPrefix string) string {
	timeStr := strconv.Itoa(int(time.Now().UnixNano()))
	timeStr = timeStr[:3] + timeStr[len(timeStr)-5:]
	userFlag := GetMD5(userPrefix)
	userFlag = userFlag[:2] + userFlag[8:10] + userFlag[12:14] + userFlag[18:20]
	machineGuid, err := getMachineGuid()
	if err != nil {
		logrus.Errorf("getMachineGuid error: %v", err)
	}
	return machineGuid[:8] + "-" + timeStr + "-" + userFlag + "-" + GetRandStr(8)
}

func GetRandStr(len int) string {
	len = len % 9
	return strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(math.Pow(10, float64(len))))))
}

func getMachineGuid() (string, error) {
	// there has been reports of issues on 32bit using golang.org/x/sys/windows/registry, see https://github.com/shirou/gopsutil/pull/312#issuecomment-277422612
	// for rationale of using windows.RegOpenKeyEx/RegQueryValueEx instead of registry.OpenKey/GetStringValue
	var h windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
	if err != nil {
		return "", err
	}
	defer windows.RegCloseKey(h)

	const windowsRegBufLen = 74 // len(`{`) + len(`abcdefgh-1234-456789012-123345456671` * 2) + len(`}`) // 2 == bytes/UTF16
	const uuidLen = 36

	var regBuf [windowsRegBufLen]uint16
	bufLen := uint32(windowsRegBufLen)
	var valType uint32
	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
	if err != nil {
		return "", err
	}

	hostID := windows.UTF16ToString(regBuf[:])
	hostIDLen := len(hostID)
	if hostIDLen != uuidLen {
		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
	}

	return hostID, nil
}

// InArrayString 判断 stub 是否在target
func InArrayString(stub string, target []string) int {
	for i, v := range target {
		if v == stub {
			return i
		}
	}
	return -1
}
