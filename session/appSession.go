package mayaSession

import (
    "maya/entity/buffer"
)

type termSession struct {
    mouseBuff mayaBuffer.mouseBuffer
    termBuff mayaBuffer.windowBuffer
}
