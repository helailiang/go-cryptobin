package gost

import "math/big"

var (
    CurveGostR34102001ParamSetcc func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xC7,
            }),
            BytesToBigint([]byte{
                0x5f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
                0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
                0x60, 0x61, 0x17, 0xa2, 0xf4, 0xbd, 0xe4, 0x28,
                0xb7, 0x45, 0x8a, 0x54, 0xb6, 0xe8, 0x7b, 0x85,
            }),
            BytesToBigint([]byte{
                0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc4,
            }),
            BytesToBigint([]byte{
                0x2d, 0x06, 0xB4, 0x26, 0x5e, 0xbc, 0x74, 0x9f,
                0xf7, 0xd0, 0xf1, 0xf1, 0xf8, 0x82, 0x32, 0xe8,
                0x16, 0x32, 0xe9, 0x08, 0x8f, 0xd4, 0x4b, 0x77,
                0x87, 0xd5, 0xe4, 0x07, 0xe9, 0x55, 0x08, 0x0c,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
            }),
            BytesToBigint([]byte{
                0xa2, 0x0e, 0x03, 0x4b, 0xf8, 0x81, 0x3e, 0xf5,
                0xc1, 0x8d, 0x01, 0x10, 0x5e, 0x72, 0x6a, 0x17,
                0xeb, 0x24, 0x8b, 0x26, 0x4a, 0xe9, 0x70, 0x6f,
                0x44, 0x0b, 0xed, 0xc8, 0xcc, 0xb6, 0xb2, 0x2c,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "GostR34102001ParamSetcc"
        return curve
    }

    // id-GostR3410-2001-TestParamSet
    CurveIdGostR34102001TestParamSet func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x31,
            }),
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
                0x50, 0xFE, 0x8A, 0x18, 0x92, 0x97, 0x61, 0x54,
                0xC5, 0x9C, 0xFC, 0x19, 0x3A, 0xCC, 0xF5, 0xB3,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
            }),
            BytesToBigint([]byte{
                0x5F, 0xBF, 0xF4, 0x98, 0xAA, 0x93, 0x8C, 0xE7,
                0x39, 0xB8, 0xE0, 0x22, 0xFB, 0xAF, 0xEF, 0x40,
                0x56, 0x3F, 0x6E, 0x6A, 0x34, 0x72, 0xFC, 0x2A,
                0x51, 0x4C, 0x0C, 0xE9, 0xDA, 0xE2, 0x3B, 0x7E,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
            }),
            BytesToBigint([]byte{
                0x08, 0xE2, 0xA8, 0xA0, 0xE6, 0x51, 0x47, 0xD4,
                0xBD, 0x63, 0x16, 0x03, 0x0E, 0x16, 0xD1, 0x9C,
                0x85, 0xC9, 0x7F, 0x0A, 0x9C, 0xA2, 0x67, 0x12,
                0x2B, 0x96, 0xAB, 0xBC, 0xEA, 0x7E, 0x8F, 0xC8,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-GostR3410-2001-TestParamSet"
        return curve
    }

    // id-tc26-gost-3410-12-256-paramSetA
    CurveIdtc26gost341012256paramSetA func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0x97,
            }),
            BytesToBigint([]byte{
                0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x0F, 0xD8, 0xCD, 0xDF, 0xC8, 0x7B, 0x66, 0x35,
                0xC1, 0x15, 0xAF, 0x55, 0x6C, 0x36, 0x0C, 0x67,
            }),
            BytesToBigint([]byte{
                0xC2, 0x17, 0x3F, 0x15, 0x13, 0x98, 0x16, 0x73,
                0xAF, 0x48, 0x92, 0xC2, 0x30, 0x35, 0xA2, 0x7C,
                0xE2, 0x5E, 0x20, 0x13, 0xBF, 0x95, 0xAA, 0x33,
                0xB2, 0x2C, 0x65, 0x6F, 0x27, 0x7E, 0x73, 0x35,
            }),
            BytesToBigint([]byte{
                0x29, 0x5F, 0x9B, 0xAE, 0x74, 0x28, 0xED, 0x9C,
                0xCC, 0x20, 0xE7, 0xC3, 0x59, 0xA9, 0xD4, 0x1A,
                0x22, 0xFC, 0xCD, 0x91, 0x08, 0xE1, 0x7B, 0xF7,
                0xBA, 0x93, 0x37, 0xA6, 0xF8, 0xAE, 0x95, 0x13,
            }),
            BytesToBigint([]byte{
                0x91, 0xE3, 0x84, 0x43, 0xA5, 0xE8, 0x2C, 0x0D,
                0x88, 0x09, 0x23, 0x42, 0x57, 0x12, 0xB2, 0xBB,
                0x65, 0x8B, 0x91, 0x96, 0x93, 0x2E, 0x02, 0xC7,
                0x8B, 0x25, 0x82, 0xFE, 0x74, 0x2D, 0xAA, 0x28,
            }),
            BytesToBigint([]byte{
                0x32, 0x87, 0x94, 0x23, 0xAB, 0x1A, 0x03, 0x75,
                0x89, 0x57, 0x86, 0xC4, 0xBB, 0x46, 0xE9, 0x56,
                0x5F, 0xDE, 0x0B, 0x53, 0x44, 0x76, 0x67, 0x40,
                0xAF, 0x26, 0x8A, 0xDB, 0x32, 0x32, 0x2E, 0x5C,
            }),
            bigInt1,
            BytesToBigint([]byte{
                0x06, 0x05, 0xF6, 0xB7, 0xC1, 0x83, 0xFA, 0x81,
                0x57, 0x8B, 0xC3, 0x9C, 0xFA, 0xD5, 0x18, 0x13,
                0x2B, 0x9D, 0xF6, 0x28, 0x97, 0x00, 0x9A, 0xF7,
                0xE5, 0x22, 0xC3, 0x2D, 0x6D, 0xC7, 0xBF, 0xFB,
            }),
            bigInt4,
        )

        curve.Name = "id-tc26-gost-3410-12-256-paramSetA"
        return curve
    }

    // id-tc26-gost-3410-12-256-paramSetB
    CurveIdtc26gost341012256paramSetB func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0x97,
            }),
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0x6C, 0x61, 0x10, 0x70, 0x99, 0x5A, 0xD1, 0x00,
                0x45, 0x84, 0x1B, 0x09, 0xB7, 0x61, 0xB8, 0x93,
            }),
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0x94,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa6,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
            }),
            BytesToBigint([]byte{
                0x8D, 0x91, 0xE4, 0x71, 0xE0, 0x98, 0x9C, 0xDA,
                0x27, 0xDF, 0x50, 0x5A, 0x45, 0x3F, 0x2B, 0x76,
                0x35, 0x29, 0x4F, 0x2D, 0xDF, 0x23, 0xE3, 0xB1,
                0x22, 0xAC, 0xC9, 0x9C, 0x9E, 0x9F, 0x1E, 0x14,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-256-paramSetB"
        return curve
    }

    // id-tc26-gost-3410-12-256-paramSetC
    CurveIdtc26gost341012256paramSetC func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x99,
            }),
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
                0x5F, 0x70, 0x0C, 0xFF, 0xF1, 0xA6, 0x24, 0xE5,
                0xE4, 0x97, 0x16, 0x1B, 0xCC, 0x8A, 0x19, 0x8F,
            }),
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x96,
            }),
            BytesToBigint([]byte{
                0x3E, 0x1A, 0xF4, 0x19, 0xA2, 0x69, 0xA5, 0xF8,
                0x66, 0xA7, 0xD3, 0xC2, 0x5C, 0x3D, 0xF8, 0x0A,
                0xE9, 0x79, 0x25, 0x93, 0x73, 0xFF, 0x2B, 0x18,
                0x2F, 0x49, 0xD4, 0xCE, 0x7E, 0x1B, 0xBC, 0x8B,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
            }),
            BytesToBigint([]byte{
                0x3F, 0xA8, 0x12, 0x43, 0x59, 0xF9, 0x66, 0x80,
                0xB8, 0x3D, 0x1C, 0x3E, 0xB2, 0xC0, 0x70, 0xE5,
                0xC5, 0x45, 0xC9, 0x85, 0x8D, 0x03, 0xEC, 0xFB,
                0x74, 0x4B, 0xF8, 0xD7, 0x17, 0x71, 0x7E, 0xFC,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-256-paramSetC"
        return curve
    }

    // id-tc26-gost-3410-12-256-paramSetD
    CurveIdtc26gost341012256paramSetD func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0x9B, 0x9F, 0x60, 0x5F, 0x5A, 0x85, 0x81, 0x07,
                0xAB, 0x1E, 0xC8, 0x5E, 0x6B, 0x41, 0xC8, 0xAA,
                0xCF, 0x84, 0x6E, 0x86, 0x78, 0x90, 0x51, 0xD3,
                0x79, 0x98, 0xF7, 0xB9, 0x02, 0x2D, 0x75, 0x9B,
            }),
            BytesToBigint([]byte{
                0x9B, 0x9F, 0x60, 0x5F, 0x5A, 0x85, 0x81, 0x07,
                0xAB, 0x1E, 0xC8, 0x5E, 0x6B, 0x41, 0xC8, 0xAA,
                0x58, 0x2C, 0xA3, 0x51, 0x1E, 0xDD, 0xFB, 0x74,
                0xF0, 0x2F, 0x3A, 0x65, 0x98, 0x98, 0x0B, 0xB9,
            }),
            BytesToBigint([]byte{
                0x9B, 0x9F, 0x60, 0x5F, 0x5A, 0x85, 0x81, 0x07,
                0xAB, 0x1E, 0xC8, 0x5E, 0x6B, 0x41, 0xC8, 0xAA,
                0xCF, 0x84, 0x6E, 0x86, 0x78, 0x90, 0x51, 0xD3,
                0x79, 0x98, 0xF7, 0xB9, 0x02, 0x2D, 0x75, 0x98,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x5a,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            }),
            BytesToBigint([]byte{
                0x41, 0xEC, 0xE5, 0x57, 0x43, 0x71, 0x1A, 0x8C,
                0x3C, 0xBF, 0x37, 0x83, 0xCD, 0x08, 0xC0, 0xEE,
                0x4D, 0x4D, 0xC4, 0x40, 0xD4, 0x64, 0x1A, 0x8F,
                0x36, 0x6E, 0x55, 0x0D, 0xFD, 0xB3, 0xBB, 0x67,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-256-paramSetD"
        return curve
    }

    // id-tc26-gost-3410-12-512-paramSetTest
    CurveIdtc26gost341012512paramSetTest func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
                0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
                0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
                0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
                0xF1, 0xD8, 0x52, 0x74, 0x1A, 0xF4, 0x70, 0x4A,
                0x04, 0x58, 0x04, 0x7E, 0x80, 0xE4, 0x54, 0x6D,
                0x35, 0xB8, 0x33, 0x6F, 0xAC, 0x22, 0x4D, 0xD8,
                0x16, 0x64, 0xBB, 0xF5, 0x28, 0xBE, 0x63, 0x73,
            }),
            BytesToBigint([]byte{
                0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
                0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
                0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
                0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
                0xA8, 0x2F, 0x2D, 0x7E, 0xCB, 0x1D, 0xBA, 0xC7,
                0x19, 0x90, 0x5C, 0x5E, 0xEC, 0xC4, 0x23, 0xF1,
                0xD8, 0x6E, 0x25, 0xED, 0xBE, 0x23, 0xC5, 0x95,
                0xD6, 0x44, 0xAA, 0xF1, 0x87, 0xE6, 0xE6, 0xDF,
            }),
            big.NewInt(7),
            BytesToBigint([]byte{
                0x1C, 0xFF, 0x08, 0x06, 0xA3, 0x11, 0x16, 0xDA,
                0x29, 0xD8, 0xCF, 0xA5, 0x4E, 0x57, 0xEB, 0x74,
                0x8B, 0xC5, 0xF3, 0x77, 0xE4, 0x94, 0x00, 0xFD,
                0xD7, 0x88, 0xB6, 0x49, 0xEC, 0xA1, 0xAC, 0x43,
                0x61, 0x83, 0x40, 0x13, 0xB2, 0xAD, 0x73, 0x22,
                0x48, 0x0A, 0x89, 0xCA, 0x58, 0xE0, 0xCF, 0x74,
                0xBC, 0x9E, 0x54, 0x0C, 0x2A, 0xDD, 0x68, 0x97,
                0xFA, 0xD0, 0xA3, 0x08, 0x4F, 0x30, 0x2A, 0xDC,
            }),
            BytesToBigint([]byte{
                0x24, 0xD1, 0x9C, 0xC6, 0x45, 0x72, 0xEE, 0x30,
                0xF3, 0x96, 0xBF, 0x6E, 0xBB, 0xFD, 0x7A, 0x6C,
                0x52, 0x13, 0xB3, 0xB3, 0xD7, 0x05, 0x7C, 0xC8,
                0x25, 0xF9, 0x10, 0x93, 0xA6, 0x8C, 0xD7, 0x62,
                0xFD, 0x60, 0x61, 0x12, 0x62, 0xCD, 0x83, 0x8D,
                0xC6, 0xB6, 0x0A, 0xA7, 0xEE, 0xE8, 0x04, 0xE2,
                0x8B, 0xC8, 0x49, 0x97, 0x7F, 0xAC, 0x33, 0xB4,
                0xB5, 0x30, 0xF1, 0xB1, 0x20, 0x24, 0x8A, 0x9A,
            }),
            BytesToBigint([]byte{
                0x2B, 0xB3, 0x12, 0xA4, 0x3B, 0xD2, 0xCE, 0x6E,
                0x0D, 0x02, 0x06, 0x13, 0xC8, 0x57, 0xAC, 0xDD,
                0xCF, 0xBF, 0x06, 0x1E, 0x91, 0xE5, 0xF2, 0xC3,
                0xF3, 0x24, 0x47, 0xC2, 0x59, 0xF3, 0x9B, 0x2C,
                0x83, 0xAB, 0x15, 0x6D, 0x77, 0xF1, 0x49, 0x6B,
                0xF7, 0xEB, 0x33, 0x51, 0xE1, 0xEE, 0x4E, 0x43,
                0xDC, 0x1A, 0x18, 0xB9, 0x1B, 0x24, 0x64, 0x0B,
                0x6D, 0xBB, 0x92, 0xCB, 0x1A, 0xDD, 0x37, 0x1E,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-512-paramSetTest"
        return curve
    }

    // id-tc26-gost-3410-12-512-paramSetA
    CurveIdtc26gost341012512paramSetA func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0xC7,
            }),
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0x27, 0xE6, 0x95, 0x32, 0xF4, 0x8D, 0x89, 0x11,
                0x6F, 0xF2, 0x2B, 0x8D, 0x4E, 0x05, 0x60, 0x60,
                0x9B, 0x4B, 0x38, 0xAB, 0xFA, 0xD2, 0xB8, 0x5D,
                0xCA, 0xCD, 0xB1, 0x41, 0x1F, 0x10, 0xB2, 0x75,
            }),
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0xC4,
            }),
            BytesToBigint([]byte{
                0xE8, 0xC2, 0x50, 0x5D, 0xED, 0xFC, 0x86, 0xDD,
                0xC1, 0xBD, 0x0B, 0x2B, 0x66, 0x67, 0xF1, 0xDA,
                0x34, 0xB8, 0x25, 0x74, 0x76, 0x1C, 0xB0, 0xE8,
                0x79, 0xBD, 0x08, 0x1C, 0xFD, 0x0B, 0x62, 0x65,
                0xEE, 0x3C, 0xB0, 0x90, 0xF3, 0x0D, 0x27, 0x61,
                0x4C, 0xB4, 0x57, 0x40, 0x10, 0xDA, 0x90, 0xDD,
                0x86, 0x2E, 0xF9, 0xD4, 0xEB, 0xEE, 0x47, 0x61,
                0x50, 0x31, 0x90, 0x78, 0x5A, 0x71, 0xC7, 0x60,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
            }),
            BytesToBigint([]byte{
                0x75, 0x03, 0xCF, 0xE8, 0x7A, 0x83, 0x6A, 0xE3,
                0xA6, 0x1B, 0x88, 0x16, 0xE2, 0x54, 0x50, 0xE6,
                0xCE, 0x5E, 0x1C, 0x93, 0xAC, 0xF1, 0xAB, 0xC1,
                0x77, 0x80, 0x64, 0xFD, 0xCB, 0xEF, 0xA9, 0x21,
                0xDF, 0x16, 0x26, 0xBE, 0x4F, 0xD0, 0x36, 0xE9,
                0x3D, 0x75, 0xE6, 0xA5, 0x0E, 0x3A, 0x41, 0xE9,
                0x80, 0x28, 0xFE, 0x5F, 0xC2, 0x35, 0xF5, 0xB8,
                0x89, 0xA5, 0x89, 0xCB, 0x52, 0x15, 0xF2, 0xA4,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-512-paramSetA"
        return curve
    }

    // id-tc26-gost-3410-12-512-paramSetB
    CurveIdtc26gost341012512paramSetB func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6F,
            }),
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
                0x49, 0xA1, 0xEC, 0x14, 0x25, 0x65, 0xA5, 0x45,
                0xAC, 0xFD, 0xB7, 0x7B, 0xD9, 0xD4, 0x0C, 0xFA,
                0x8B, 0x99, 0x67, 0x12, 0x10, 0x1B, 0xEA, 0x0E,
                0xC6, 0x34, 0x6C, 0x54, 0x37, 0x4F, 0x25, 0xBD,
            }),
            BytesToBigint([]byte{
                0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6C,
            }),
            BytesToBigint([]byte{
                0x68, 0x7D, 0x1B, 0x45, 0x9D, 0xC8, 0x41, 0x45,
                0x7E, 0x3E, 0x06, 0xCF, 0x6F, 0x5E, 0x25, 0x17,
                0xB9, 0x7C, 0x7D, 0x61, 0x4A, 0xF1, 0x38, 0xBC,
                0xBF, 0x85, 0xDC, 0x80, 0x6C, 0x4B, 0x28, 0x9F,
                0x3E, 0x96, 0x5D, 0x2D, 0xB1, 0x41, 0x6D, 0x21,
                0x7F, 0x8B, 0x27, 0x6F, 0xAD, 0x1A, 0xB6, 0x9C,
                0x50, 0xF7, 0x8B, 0xEE, 0x1F, 0xA3, 0x10, 0x6E,
                0xFB, 0x8C, 0xCB, 0xC7, 0xC5, 0x14, 0x01, 0x16,
            }),
            BytesToBigint([]byte{
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
            }),
            BytesToBigint([]byte{
                0x1A, 0x8F, 0x7E, 0xDA, 0x38, 0x9B, 0x09, 0x4C,
                0x2C, 0x07, 0x1E, 0x36, 0x47, 0xA8, 0x94, 0x0F,
                0x3C, 0x12, 0x3B, 0x69, 0x75, 0x78, 0xC2, 0x13,
                0xBE, 0x6D, 0xD9, 0xE6, 0xC8, 0xEC, 0x73, 0x35,
                0xDC, 0xB2, 0x28, 0xFD, 0x1E, 0xDF, 0x4A, 0x39,
                0x15, 0x2C, 0xBC, 0xAA, 0xF8, 0xC0, 0x39, 0x88,
                0x28, 0x04, 0x10, 0x55, 0xF9, 0x4C, 0xEE, 0xEC,
                0x7E, 0x21, 0x34, 0x07, 0x80, 0xFE, 0x41, 0xBD,
            }),
            nil,
            nil,
            nil,
        )

        curve.Name = "id-tc26-gost-3410-12-512-paramSetB"
        return curve
    }

    // id-tc26-gost-3410-12-512-paramSetC
    CurveIdtc26gost341012512paramSetC func() *Curve = func() *Curve {
        curve, _ := NewCurve(
            BytesToBigint([]byte{
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFD, 0xC7,
            }),
            BytesToBigint([]byte{
                0x3F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xC9, 0x8C, 0xDB, 0xA4, 0x65, 0x06, 0xAB, 0x00,
                0x4C, 0x33, 0xA9, 0xFF, 0x51, 0x47, 0x50, 0x2C,
                0xC8, 0xED, 0xA9, 0xE7, 0xA7, 0x69, 0xA1, 0x26,
                0x94, 0x62, 0x3C, 0xEF, 0x47, 0xF0, 0x23, 0xED,
            }),
            BytesToBigint([]byte{
                0xDC, 0x92, 0x03, 0xE5, 0x14, 0xA7, 0x21, 0x87,
                0x54, 0x85, 0xA5, 0x29, 0xD2, 0xC7, 0x22, 0xFB,
                0x18, 0x7B, 0xC8, 0x98, 0x0E, 0xB8, 0x66, 0x64,
                0x4D, 0xE4, 0x1C, 0x68, 0xE1, 0x43, 0x06, 0x45,
                0x46, 0xE8, 0x61, 0xC0, 0xE2, 0xC9, 0xED, 0xD9,
                0x2A, 0xDE, 0x71, 0xF4, 0x6F, 0xCF, 0x50, 0xFF,
                0x2A, 0xD9, 0x7F, 0x95, 0x1F, 0xDA, 0x9F, 0x2A,
                0x2E, 0xB6, 0x54, 0x6F, 0x39, 0x68, 0x9B, 0xD3,
            }),
            BytesToBigint([]byte{
                0xB4, 0xC4, 0xEE, 0x28, 0xCE, 0xBC, 0x6C, 0x2C,
                0x8A, 0xC1, 0x29, 0x52, 0xCF, 0x37, 0xF1, 0x6A,
                0xC7, 0xEF, 0xB6, 0xA9, 0xF6, 0x9F, 0x4B, 0x57,
                0xFF, 0xDA, 0x2E, 0x4F, 0x0D, 0xE5, 0xAD, 0xE0,
                0x38, 0xCB, 0xC2, 0xFF, 0xF7, 0x19, 0xD2, 0xC1,
                0x8D, 0xE0, 0x28, 0x4B, 0x8B, 0xFE, 0xF3, 0xB5,
                0x2B, 0x8C, 0xC7, 0xA5, 0xF5, 0xBF, 0x0A, 0x3C,
                0x8D, 0x23, 0x19, 0xA5, 0x31, 0x25, 0x57, 0xE1,
            }),
            BytesToBigint([]byte{
                0xE2, 0xE3, 0x1E, 0xDF, 0xC2, 0x3D, 0xE7, 0xBD,
                0xEB, 0xE2, 0x41, 0xCE, 0x59, 0x3E, 0xF5, 0xDE,
                0x22, 0x95, 0xB7, 0xA9, 0xCB, 0xAE, 0xF0, 0x21,
                0xD3, 0x85, 0xF7, 0x07, 0x4C, 0xEA, 0x04, 0x3A,
                0xA2, 0x72, 0x72, 0xA7, 0xAE, 0x60, 0x2B, 0xF2,
                0xA7, 0xB9, 0x03, 0x3D, 0xB9, 0xED, 0x36, 0x10,
                0xC6, 0xFB, 0x85, 0x48, 0x7E, 0xAE, 0x97, 0xAA,
                0xC5, 0xBC, 0x79, 0x28, 0xC1, 0x95, 0x01, 0x48,
            }),
            BytesToBigint([]byte{
                0xF5, 0xCE, 0x40, 0xD9, 0x5B, 0x5E, 0xB8, 0x99,
                0xAB, 0xBC, 0xCF, 0xF5, 0x91, 0x1C, 0xB8, 0x57,
                0x79, 0x39, 0x80, 0x4D, 0x65, 0x27, 0x37, 0x8B,
                0x8C, 0x10, 0x8C, 0x3D, 0x20, 0x90, 0xFF, 0x9B,
                0xE1, 0x8E, 0x2D, 0x33, 0xE3, 0x02, 0x1E, 0xD2,
                0xEF, 0x32, 0xD8, 0x58, 0x22, 0x42, 0x3B, 0x63,
                0x04, 0xF7, 0x26, 0xAA, 0x85, 0x4B, 0xAE, 0x07,
                0xD0, 0x39, 0x6E, 0x9A, 0x9A, 0xDD, 0xC4, 0x0F,
            }),
            bigInt1,
            BytesToBigint([]byte{
                0x9E, 0x4F, 0x5D, 0x8C, 0x01, 0x7D, 0x8D, 0x9F,
                0x13, 0xA5, 0xCF, 0x3C, 0xDF, 0x5B, 0xFE, 0x4D,
                0xAB, 0x40, 0x2D, 0x54, 0x19, 0x8E, 0x31, 0xEB,
                0xDE, 0x28, 0xA0, 0x62, 0x10, 0x50, 0x43, 0x9C,
                0xA6, 0xB3, 0x9E, 0x0A, 0x51, 0x5C, 0x06, 0xB3,
                0x04, 0xE2, 0xCE, 0x43, 0xE7, 0x9E, 0x36, 0x9E,
                0x91, 0xA0, 0xCF, 0xC2, 0xBC, 0x2A, 0x22, 0xB4,
                0xCA, 0x30, 0x2D, 0xBB, 0x33, 0xEE, 0x75, 0x50,
            }),
            bigInt4,
        )

        curve.Name = "id-tc26-gost-3410-12-512-paramSetC"
        return curve
    }

    // Aliases

    // id-GostR3410-2001-CryptoPro-A-ParamSet
    CurveIdGostR34102001CryptoProAParamSet func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetB()
        c.Name = "id-GostR3410-2001-CryptoPro-A-ParamSet"
        return c
    }

    // id-GostR3410-2001-CryptoPro-B-ParamSet
    CurveIdGostR34102001CryptoProBParamSet func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetC()
        c.Name = "id-GostR3410-2001-CryptoPro-B-ParamSet"
        return c
    }

    // id-GostR3410-2001-CryptoPro-C-ParamSet
    CurveIdGostR34102001CryptoProCParamSet func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetD()
        c.Name = "id-GostR3410-2001-CryptoPro-C-ParamSet"
        return c
    }

    // id-GostR3410-2001-CryptoPro-XchA-ParamSet
    CurveIdGostR34102001CryptoProXchAParamSet func() *Curve = func() *Curve {
        c := CurveIdGostR34102001CryptoProAParamSet()
        c.Name = "id-GostR3410-2001-CryptoPro-XchA-ParamSet"
        return c
    }

    // id-GostR3410-2001-CryptoPro-XchB-ParamSet
    CurveIdGostR34102001CryptoProXchBParamSet func() *Curve = func() *Curve {
        c := CurveIdGostR34102001CryptoProCParamSet()
        c.Name = "id-GostR3410-2001-CryptoPro-XchB-ParamSet"
        return c
    }

    // id-tc26-gost-3410-2012-256-paramSetA
    CurveIdtc26gost34102012256paramSetA func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetA()
        c.Name = "id-tc26-gost-3410-2012-256-paramSetA"
        return c
    }

    // id-tc26-gost-3410-2012-256-paramSetB
    CurveIdtc26gost34102012256paramSetB func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetB()
        c.Name = "id-tc26-gost-3410-2012-256-paramSetB"
        return c
    }

    // id-tc26-gost-3410-2012-256-paramSetC
    CurveIdtc26gost34102012256paramSetC func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetC()
        c.Name = "id-tc26-gost-3410-2012-256-paramSetC"
        return c
    }

    // id-tc26-gost-3410-2012-256-paramSetD
    CurveIdtc26gost34102012256paramSetD func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012256paramSetD()
        c.Name = "id-tc26-gost-3410-2012-256-paramSetD"
        return c
    }

    // id-tc26-gost-3410-2012-512-paramSetTest
    CurveIdtc26gost34102012512paramSetTest func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012512paramSetTest()
        c.Name = "id-tc26-gost-3410-2012-512-paramSetTest"
        return c
    }

    // id-tc26-gost-3410-2012-512-paramSetA
    CurveIdtc26gost34102012512paramSetA func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012512paramSetA()
        c.Name = "id-tc26-gost-3410-2012-512-paramSetA"
        return c
    }

    // id-tc26-gost-3410-2012-512-paramSetB
    CurveIdtc26gost34102012512paramSetB func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012512paramSetB()
        c.Name = "id-tc26-gost-3410-2012-512-paramSetB"
        return c
    }

    // id-tc26-gost-3410-2012-512-paramSetC
    CurveIdtc26gost34102012512paramSetC func() *Curve = func() *Curve {
        c := CurveIdtc26gost341012512paramSetC()
        c.Name = "id-tc26-gost-3410-2012-512-paramSetC"
        return c
    }

    CurveDefault = CurveIdtc26gost341012256paramSetB
)
