package keyboard

import (
	"encoding/hex"
	"errors"
	"reflect"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

type keyboard struct {
}

type Key int

func lazyAddr(pAddr *uintptr, lib *windows.LazyDLL, procName string) uintptr {
	addr := atomic.LoadUintptr(pAddr)
	if addr == 0 {
		addr = lib.NewProc(procName).Addr()
		atomic.StoreUintptr(pAddr, addr)
	}
	return addr
}

var (
	libUser32   = windows.NewLazySystemDLL("user32.dll")
)

var (
	pGetKeyboardLayout uintptr
	pVkKeyScanExA uintptr
)

const (
	KEY_CANCEL Key = iota
	KEY_XBUTTON1
	KEY_XBUTTON2
	KEY_BACK
	KEY_TAB
	KEY_CLEAR
	KEY_RETURN
	KEY_SHIFT
	KEY_CONTROL
	KEY_MENU
	KEY_PAUSE
	KEY_CAPITAL
	KEY_KANA
	KEY_HANGUL
	KEY_IME_ON
	KEY_JUNJA
	KEY_FINAL
	KEY_HANJA
	KEY_KANJI
	KEY_IME_OFF
	KEY_ESCAPE
	KEY_CONVERT
	KEY_NONCONVERT
	KEY_ACCEPT
	KEY_MODECHANGE
	KEY_SPACE
	KEY_PRIOR
	KEY_NEXT
	KEY_END
	KEY_HOME
	KEY_LEFT
	KEY_UP
	KEY_RIGHT
	KEY_DOWN
	KEY_SELECT
	KEY_PRINT
	KEY_EXECUTE
	KEY_SNAPSHOT
	KEY_INSERT
	KEY_DELETE
	KEY_HELP
	KEY_0
	KEY_1
	KEY_2
	KEY_3
	KEY_4
	KEY_5
	KEY_6
	KEY_7
	KEY_8
	KEY_9
	KEY_A
	KEY_B
	KEY_C
	KEY_D
	KEY_E
	KEY_F
	KEY_G
	KEY_H
	KEY_I
	KEY_J
	KEY_K
	KEY_L
	KEY_M
	KEY_N
	KEY_O
	KEY_P
	KEY_Q
	KEY_R
	KEY_S
	KEY_T
	KEY_U
	KEY_V
	KEY_W
	KEY_X
	KEY_Y
	KEY_Z
	KEY_LWIN
	KEY_RWIN
	KEY_APPS
	KEY_SLEEP
	KEY_NUMPAD0
	KEY_NUMPAD1
	KEY_NUMPAD2
	KEY_NUMPAD3
	KEY_NUMPAD4
	KEY_NUMPAD5
	KEY_NUMPAD6
	KEY_NUMPAD7
	KEY_NUMPAD8
	KEY_NUMPAD9
	KEY_MULTIPLY
	KEY_ADD
	KEY_SEPARATOR
	KEY_SUBTRACT
	KEY_DECIMAL
	KEY_DIVIDE
	KEY_F1
	KEY_F2
	KEY_F3
	KEY_F4
	KEY_F5
	KEY_F6
	KEY_F7
	KEY_F8
	KEY_F9
	KEY_F10
	KEY_F11
	KEY_F12
	KEY_F13
	KEY_F14
	KEY_F15
	KEY_F16
	KEY_F17
	KEY_F18
	KEY_F19
	KEY_F20
	KEY_F21
	KEY_F22
	KEY_F23
	KEY_F24
	KEY_NUMLOCK
	KEY_SCROLL
	KEY_LSHIFT
	KEY_RSHIFT
	KEY_LCONTROL
	KEY_RCONTROL
	KEY_LMENU
	KEY_RMENU
	KEY_BROWSER_BACK
	KEY_BROWSER_FORWARD
	KEY_BROWSER_REFRESH
	KEY_BROWSER_STOP
	KEY_BROWSER_SEARCH
	KEY_BROWSER_FAVORITES
	KEY_BROWSER_HOME
	KEY_VOLUME_MUTE
	KEY_VOLUME_DOWN
	KEY_VOLUME_UP
	KEY_MEDIA_NEXT_TRACK
	KEY_MEDIA_PREV_TRACK
	KEY_MEDIA_STOP
	KEY_MEDIA_PLAY_PAUSE
	KEY_LAUNCH_MAIL
	KEY_LAUNCH_MEDIA_SELECT
	KEY_LAUNCH_APP1
	KEY_LAUNCH_APP2
	KEY_OEM_1
	KEY_OEM_PLUS
	KEY_OEM_COMMA
	KEY_OEM_MINUS
	KEY_OEM_PERIOD
	KEY_OEM_2
	KEY_OEM_3
	KEY_OEM_4
	KEY_OEM_5
	KEY_OEM_6
	KEY_OEM_7
	KEY_OEM_8
	KEY_OEM_102
	KEY_PROCESSKEY
	KEY_ATTN
	KEY_CRSEL
	KEY_EXSEL
	KEY_EREOF
	KEY_PLAY
	KEY_ZOOM
	KEY_NONAME
	KEY_PA1
	KEY_OEM_CLEAR
)



func getKeyOrCharHex(key interface{}) (uint16, error) {
	switch key.(type) {
		case Key:
			return getKeyHex(reflect.ValueOf(key).Interface().(Key)), nil
		case int32:
			kbl := getKeyboardLayout(0)
			return uint16(vkKeyScanExA(uint8(reflect.ValueOf(key).Interface().(int32)), kbl)), nil
		default:
			return 0, errors.New("getKeyOrCharHex() can only accept characters and Keys")
	}
}

func getKeyHex(key Key) uint16 {
	var res uint16 = 0;
  
	switch key {
		case KEY_CANCEL:
			res = 0x03;
		  

		case KEY_XBUTTON1:
			res = 0x05;
		  

		case KEY_XBUTTON2:
			res = 0x06;
		  

		case KEY_BACK:
			res = 0x08;
		  

		case KEY_TAB:
			res = 0x09;
		  

		case KEY_CLEAR:
			res = 0x0C;
		  

		case KEY_RETURN:
			res = 0x0D;
		  

		case KEY_SHIFT:
			res = 0x10;
		  

		case KEY_CONTROL:
			res = 0x11;
		  

		case KEY_MENU:
			res = 0x12;
		  

		case KEY_PAUSE:
			res = 0x13;
		  

		case KEY_CAPITAL:
			res = 0x14;
		  

		case KEY_KANA:
			res = 0x15;
		  

		case KEY_HANGUL:
			res = 0x15;
		  

		case KEY_IME_ON:
			res = 0x16;
		  

		case KEY_JUNJA:
			res = 0x17;
		  

		case KEY_FINAL:
			res = 0x18;
		  

		case KEY_HANJA:
			res = 0x19;
		  

		case KEY_KANJI:
			res = 0x19;
		  

		case KEY_IME_OFF:
			res = 0x1A;
		  

		case KEY_ESCAPE:
			res = 0x1B;
		  

		case KEY_CONVERT:
			res = 0x1C;
		  

		case KEY_NONCONVERT:
			res = 0x1D;
		  

		case KEY_ACCEPT:
			res = 0x1E;
		  

		case KEY_MODECHANGE:
			res = 0x1F;
		  

		case KEY_SPACE:
			res = 0x20;
		  

		case KEY_PRIOR:
			res = 0x21;
		  

		case KEY_NEXT:
			res = 0x22;
		  

		case KEY_END:
			res = 0x23;
		  

		case KEY_HOME:
			res = 0x24;
		  

		case KEY_LEFT:
			res = 0x25;
		  

		case KEY_UP:
			res = 0x26;
		  

		case KEY_RIGHT:
			res = 0x27;
		  

		case KEY_DOWN:
			res = 0x28;
		  

		case KEY_SELECT:
			res = 0x29;
		  

		case KEY_PRINT:
			res = 0x2A;
		  

		case KEY_EXECUTE:
			res = 0x2B;
		  

		case KEY_SNAPSHOT:
			res = 0x2C;
		  

		case KEY_INSERT:
			res = 0x2D;
		  

		case KEY_DELETE:
			res = 0x2E;
		  

		case KEY_HELP:
			res = 0x2F;
		  

		case KEY_0:
			res = 0x30;
		  

		case KEY_1:
			res = 0x31;
		  

		case KEY_2:
			res = 0x32;
		  

		case KEY_3:
			res = 0x33;
		  

		case KEY_4:
			res = 0x34;
		  

		case KEY_5:
			res = 0x35;
		  

		case KEY_6:
			res = 0x36;
		  

		case KEY_7:
			res = 0x37;
		  

		case KEY_8:
			res = 0x38;
		  

		case KEY_9:
			res = 0x39;
		  

		case KEY_A:
			res = 0x41;
		  

		case KEY_B:
			res = 0x42;
		  

		case KEY_C:
			res = 0x43;
		  

		case KEY_D:
			res = 0x44;
		  

		case KEY_E:
			res = 0x45;
		  

		case KEY_F:
			res = 0x46;
		  

		case KEY_G:
			res = 0x47;
		  

		case KEY_H:
			res = 0x48;
		  

		case KEY_I:
			res = 0x49;
		  

		case KEY_J:
			res = 0x4A;
		  

		case KEY_K:
			res = 0x4B;
		  

		case KEY_L:
			res = 0x4C;
		  

		case KEY_M:
			res = 0x4D;
		  

		case KEY_N:
			res = 0x4E;
		  

		case KEY_O:
			res = 0x4F;
		  

		case KEY_P:
			res = 0x50;
		  

		case KEY_Q:
			res = 0x51;
		  

		case KEY_R:
			res = 0x52;
		  

		case KEY_S:
			res = 0x53;
		  

		case KEY_T:
			res = 0x54;
		  

		case KEY_U:
			res = 0x55;
		  

		case KEY_V:
			res = 0x56;
		  

		case KEY_W:
			res = 0x57;
		  

		case KEY_X:
			res = 0x58;
		  

		case KEY_Y:
			res = 0x59;
		  

		case KEY_Z:
			res = 0x5A;
		  

		case KEY_LWIN:
			res = 0x5B;
		  

		case KEY_RWIN:
			res = 0x5C;
		  

		case KEY_APPS:
			res = 0x5D;
		  

		case KEY_SLEEP:
			res = 0x5F;
		  

		case KEY_NUMPAD0:
			res = 0x60;
		  

		case KEY_NUMPAD1:
			res = 0x61;
		  

		case KEY_NUMPAD2:
			res = 0x62;
		  

		case KEY_NUMPAD3:
			res = 0x63;
		  

		case KEY_NUMPAD4:
			res = 0x64;
		  

		case KEY_NUMPAD5:
			res = 0x65;
		  

		case KEY_NUMPAD6:
			res = 0x66;
		  

		case KEY_NUMPAD7:
			res = 0x67;
		  

		case KEY_NUMPAD8:
			res = 0x68;
		  

		case KEY_NUMPAD9:
			res = 0x69;
		  

		case KEY_MULTIPLY:
			res = 0x6A;
		  

		case KEY_ADD:
			res = 0x6B;
		  

		case KEY_SEPARATOR:
			res = 0x6C;
		  

		case KEY_SUBTRACT:
			res = 0x6D;
		  

		case KEY_DECIMAL:
			res = 0x6E;
		  

		case KEY_DIVIDE:
			res = 0x6F;
		  

		case KEY_F1:
			res = 0x70;
		  

		case KEY_F2:
			res = 0x71;
		  

		case KEY_F3:
			res = 0x72;
		  

		case KEY_F4:
			res = 0x73;
		  

		case KEY_F5:
			res = 0x74;
		  

		case KEY_F6:
			res = 0x75;
		  

		case KEY_F7:
			res = 0x76;
		  

		case KEY_F8:
			res = 0x77;
		  

		case KEY_F9:
			res = 0x78;
		  

		case KEY_F10:
			res = 0x79;
		  

		case KEY_F11:
			res = 0x7A;
		  

		case KEY_F12:
			res = 0x7B;
		  

		case KEY_F13:
			res = 0x7C;
		  

		case KEY_F14:
			res = 0x7D;
		  

		case KEY_F15:
			res = 0x7E;
		  

		case KEY_F16:
			res = 0x7F;
		  

		case KEY_F17:
			res = 0x80;
		  

		case KEY_F18:
			res = 0x81;
		  

		case KEY_F19:
			res = 0x82;
		  

		case KEY_F20:
			res = 0x83;
		  

		case KEY_F21:
			res = 0x84;
		  

		case KEY_F22:
			res = 0x85;
		  

		case KEY_F23:
			res = 0x86;
		  

		case KEY_F24:
			res = 0x87;
		  

		case KEY_NUMLOCK:
			res = 0x90;
		  

		case KEY_SCROLL:
			res = 0x91;
		  

		case KEY_LSHIFT:
			res = 0xA0;
		  

		case KEY_RSHIFT:
			res = 0xA1;
		  

		case KEY_LCONTROL:
			res = 0xA2;
		  

		case KEY_RCONTROL:
			res = 0xA3;
		  

		case KEY_LMENU:
			res = 0xA4;
		  

		case KEY_RMENU:
			res = 0xA5;
		  

		case KEY_BROWSER_BACK:
			res = 0xA6;
		  

		case KEY_BROWSER_FORWARD:
			res = 0xA7;
		  

		case KEY_BROWSER_REFRESH:
			res = 0xA8;
		  

		case KEY_BROWSER_STOP:
			res = 0xA9;
		  

		case KEY_BROWSER_SEARCH:
			res = 0xAA;
		  

		case KEY_BROWSER_FAVORITES:
			res = 0xAB;
		  

		case KEY_BROWSER_HOME:
			res = 0xAC;
		  

		case KEY_VOLUME_MUTE:
			res = 0xAD;
		  

		case KEY_VOLUME_DOWN:
			res = 0xAE;
		  

		case KEY_VOLUME_UP:
			res = 0xAF;
		  

		case KEY_MEDIA_NEXT_TRACK:
			res = 0xB0;
		  

		case KEY_MEDIA_PREV_TRACK:
			res = 0xB1;
		  

		case KEY_MEDIA_STOP:
			res = 0xB2;
		  

		case KEY_MEDIA_PLAY_PAUSE:
			res = 0xB3;
		  

		case KEY_LAUNCH_MAIL:
			res = 0xB4;
		  

		case KEY_LAUNCH_MEDIA_SELECT:
			res = 0xB5;
		  

		case KEY_LAUNCH_APP1:
			res = 0xB6;
		  

		case KEY_LAUNCH_APP2:
			res = 0xB7;
		  

		case KEY_OEM_1:
			res = 0xBA;
		  

		case KEY_OEM_PLUS:
			res = 0xBB;
		  

		case KEY_OEM_COMMA:
			res = 0xBC;
		  

		case KEY_OEM_MINUS:
			res = 0xBD;
		  

		case KEY_OEM_PERIOD:
			res = 0xBE;
		  

		case KEY_OEM_2:
			res = 0xBF;
		  

		case KEY_OEM_3:
			res = 0xC0;
		  

		case KEY_OEM_4:
			res = 0xDB;
		  

		case KEY_OEM_5:
			res = 0xDC;
		  

		case KEY_OEM_6:
			res = 0xDD;
		  

		case KEY_OEM_7:
			res = 0xDE;
		  

		case KEY_OEM_8:
			res = 0xDF;
		  

		case KEY_OEM_102:
			res = 0xE2;
		  

		case KEY_PROCESSKEY:
			res = 0xE5;
		  

		case KEY_ATTN:
			res = 0xF6;
		  

		case KEY_CRSEL:
			res = 0xF7;
		  

		case KEY_EXSEL:
			res = 0xF8;
		  

		case KEY_EREOF:
			res = 0xF9;
		  

		case KEY_PLAY:
			res = 0xFA;
		  

		case KEY_ZOOM:
			res = 0xFB;
		  

		case KEY_NONAME:
			res = 0xFC;
		  

		case KEY_PA1:
			res = 0xFD;
		  

		case KEY_OEM_CLEAR:
			res = 0xFE;
		  

	}
  
	return res;
  }

// func GetKeyboardLayout(threadID uint32) (locale syscall.Handle) {
// 	r0, _, _ := syscall.Syscall(procGetKeyboardLayout.Addr(), 1, uintptr(threadID), 0, 0)
// 	locale = syscall.Handle(r0)
// 	return
// }
type HKL = uintptr
type CHAR = uint8

func vkKeyScanExA(ch CHAR, dwhkl HKL) int16 {
	addr := lazyAddr(&pVkKeyScanExA, libUser32, "VkKeyScanExA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(ch), dwhkl)
	return int16(ret)
}

func getKeyboardLayout(idThread uint32) HKL {
	addr := lazyAddr(&pGetKeyboardLayout, libUser32, "GetKeyboardLayout")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(idThread))
	return HKL(ret)
}



// "0061" to 97 (a)
func getIntFromHexString(h string) int {
	i, _ := strconv.ParseInt(h, 16, 32)
	return int(i)
}

// "a" to 62
func getHexStringFromRune(h rune) string {
	return hex.EncodeToString([]byte(string(h)))
}

func unicodeDown(char rune) {
	hx := getIntFromHexString(getHexStringFromRune(char))
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwFlags: win.KEYEVENTF_UNICODE,
				WScan: uint16(hx),
			},
		},
	}

    win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func unicodeUp(char rune) {
	hx := getIntFromHexString(getHexStringFromRune(char))
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwFlags: win.KEYEVENTF_UNICODE | win.KEYEVENTF_KEYUP,
				WScan: uint16(hx),
			},
		},
	}

    win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func KeyDown(hx uint16) {
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwFlags: 0,
				WVk: hx,
			},
		},
	}

    win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func KeyUp(hx uint16) {
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwFlags: win.KEYEVENTF_KEYUP,
				WVk: hx,
			},
		},
	}

    win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

// func Tap(key Key){
// 	KeyDown(key)
// 	time.Sleep(20 * time.Millisecond)
// 	KeyUp(key)
// 	time.Sleep(20 * time.Millisecond)
// }

// func Hotkey(keys ...interface{}){
//     switch reflect.TypeOf(keys).Kind() {
//     case reflect.Slice:
//         s := reflect.ValueOf(keys)
//         fmt.Println(s.Index(0))
// 		if s.Len() > 1 {
// 			Hotkey(s.Slice(1, s.Len() - 1))
// 		}
//     }

// 	// KeyDown(keys[0])
// 	// time.Sleep(20 * time.Millisecond)
// 	// if len(keys) > 1 {
// 	// 	Hotkey(keys[1:]...)
// 	// }
// 	// KeyUp(keys[0])
// }

func New() *keyboard {
	return &keyboard{}
}

func (k keyboard) Type(text string){
	for _, c := range text {
		unicodeDown(c)
		time.Sleep(20 * time.Millisecond)
		unicodeUp(c)
		time.Sleep(20 * time.Millisecond)
	}
}

func (k keyboard) Tap(key interface{}) (error) {
	switch key.(type) {
	case Key:
		val := reflect.ValueOf(key)
		hx := getKeyHex(Key(val.Int()))
		KeyDown(hx)
		time.Sleep(20 * time.Millisecond)
		KeyUp(hx)
		time.Sleep(20 * time.Millisecond)
	case int32:
		val := reflect.ValueOf(key)
		kbl := getKeyboardLayout(0)
		hx := uint16(vkKeyScanExA(uint8(val.Int()), kbl))
		
		shift_pressed := hx & 0x100 == 0x100
		ctrl_pressed := hx & 0x200 == 0x200
		alt_pressed := hx & 0x400 == 0x400
		
		if shift_pressed {
			KeyDown(getKeyHex(KEY_SHIFT))
			time.Sleep(20 * time.Millisecond)
		}
		
		if ctrl_pressed {
			KeyDown(getKeyHex(KEY_CONTROL))
			time.Sleep(20 * time.Millisecond)
		}
		
		if alt_pressed {
			KeyDown(getKeyHex(KEY_MENU))
			time.Sleep(20 * time.Millisecond)
		}

		KeyDown(hx)
		time.Sleep(20 * time.Millisecond)
		KeyUp(hx)
		time.Sleep(20 * time.Millisecond)

		if shift_pressed {
			KeyUp(getKeyHex(KEY_SHIFT))
		}

		if ctrl_pressed {
			KeyUp(getKeyHex(KEY_CONTROL))
		}

		if alt_pressed {
			KeyUp(getKeyHex(KEY_MENU))
		}
	default:
		return errors.New("Tap() can only accept characters and Keys")
	}
	return nil
}

func (k keyboard) Hotkey(keys ...interface{}) (error) {
	hx, _ := getKeyOrCharHex(keys[0])
	KeyDown(hx)
	time.Sleep(20 * time.Millisecond)
	if len(keys) > 1 {
		k.Hotkey(keys[1:]...)
	}
	KeyUp(hx)
	return nil
}
