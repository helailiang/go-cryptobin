package field

import "math/bits"

type p256Uint1 uint64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/field-crypto/pull/1006#issuecomment-892625927
type p256Int1 int64   // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/field-crypto/pull/1006#issuecomment-892625927

// The type p256NonMontgomeryDomainFieldElement is a field element NOT in the Montgomery domain.
//
// Bounds: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
type p256NonMontgomeryDomainFieldElement [4]uint64

// p256CmovznzU64 is a single-word conditional move.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xffffffffffffffff]
//   arg3: [0x0 ~> 0xffffffffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffffffffffff]
func p256CmovznzU64(out1 *uint64, arg1 p256Uint1, arg2 uint64, arg3 uint64) {
    x1 := (uint64(arg1) * 0xffffffffffffffff)
    x2 := ((x1 & arg3) | ((^x1) & arg2))
    *out1 = x2
}

// p256Mul multiplies two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func p256Mul(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
    x1 := arg1[1]
    x2 := arg1[2]
    x3 := arg1[3]
    x4 := arg1[0]
    var x5 uint64
    var x6 uint64
    x6, x5 = bits.Mul64(x4, arg2[3])
    var x7 uint64
    var x8 uint64
    x8, x7 = bits.Mul64(x4, arg2[2])
    var x9 uint64
    var x10 uint64
    x10, x9 = bits.Mul64(x4, arg2[1])
    var x11 uint64
    var x12 uint64
    x12, x11 = bits.Mul64(x4, arg2[0])
    var x13 uint64
    var x14 uint64
    x13, x14 = bits.Add64(x12, x9, uint64(0x0))
    var x15 uint64
    var x16 uint64
    x15, x16 = bits.Add64(x10, x7, uint64(p256Uint1(x14)))
    var x17 uint64
    var x18 uint64
    x17, x18 = bits.Add64(x8, x5, uint64(p256Uint1(x16)))
    x19 := (uint64(p256Uint1(x18)) + x6)
    var x20 uint64
    var x21 uint64
    x21, x20 = bits.Mul64(x11, 0xfffffffeffffffff)
    var x22 uint64
    var x23 uint64
    x23, x22 = bits.Mul64(x11, 0xffffffffffffffff)
    var x24 uint64
    var x25 uint64
    x25, x24 = bits.Mul64(x11, 0xffffffff00000000)
    var x26 uint64
    var x27 uint64
    x27, x26 = bits.Mul64(x11, 0xffffffffffffffff)
    var x28 uint64
    var x29 uint64
    x28, x29 = bits.Add64(x27, x24, uint64(0x0))
    var x30 uint64
    var x31 uint64
    x30, x31 = bits.Add64(x25, x22, uint64(p256Uint1(x29)))
    var x32 uint64
    var x33 uint64
    x32, x33 = bits.Add64(x23, x20, uint64(p256Uint1(x31)))
    x34 := (uint64(p256Uint1(x33)) + x21)
    var x36 uint64
    _, x36 = bits.Add64(x11, x26, uint64(0x0))
    var x37 uint64
    var x38 uint64
    x37, x38 = bits.Add64(x13, x28, uint64(p256Uint1(x36)))
    var x39 uint64
    var x40 uint64
    x39, x40 = bits.Add64(x15, x30, uint64(p256Uint1(x38)))
    var x41 uint64
    var x42 uint64
    x41, x42 = bits.Add64(x17, x32, uint64(p256Uint1(x40)))
    var x43 uint64
    var x44 uint64
    x43, x44 = bits.Add64(x19, x34, uint64(p256Uint1(x42)))
    var x45 uint64
    var x46 uint64
    x46, x45 = bits.Mul64(x1, arg2[3])
    var x47 uint64
    var x48 uint64
    x48, x47 = bits.Mul64(x1, arg2[2])
    var x49 uint64
    var x50 uint64
    x50, x49 = bits.Mul64(x1, arg2[1])
    var x51 uint64
    var x52 uint64
    x52, x51 = bits.Mul64(x1, arg2[0])
    var x53 uint64
    var x54 uint64
    x53, x54 = bits.Add64(x52, x49, uint64(0x0))
    var x55 uint64
    var x56 uint64
    x55, x56 = bits.Add64(x50, x47, uint64(p256Uint1(x54)))
    var x57 uint64
    var x58 uint64
    x57, x58 = bits.Add64(x48, x45, uint64(p256Uint1(x56)))
    x59 := (uint64(p256Uint1(x58)) + x46)
    var x60 uint64
    var x61 uint64
    x60, x61 = bits.Add64(x37, x51, uint64(0x0))
    var x62 uint64
    var x63 uint64
    x62, x63 = bits.Add64(x39, x53, uint64(p256Uint1(x61)))
    var x64 uint64
    var x65 uint64
    x64, x65 = bits.Add64(x41, x55, uint64(p256Uint1(x63)))
    var x66 uint64
    var x67 uint64
    x66, x67 = bits.Add64(x43, x57, uint64(p256Uint1(x65)))
    var x68 uint64
    var x69 uint64
    x68, x69 = bits.Add64(uint64(p256Uint1(x44)), x59, uint64(p256Uint1(x67)))
    var x70 uint64
    var x71 uint64
    x71, x70 = bits.Mul64(x60, 0xfffffffeffffffff)
    var x72 uint64
    var x73 uint64
    x73, x72 = bits.Mul64(x60, 0xffffffffffffffff)
    var x74 uint64
    var x75 uint64
    x75, x74 = bits.Mul64(x60, 0xffffffff00000000)
    var x76 uint64
    var x77 uint64
    x77, x76 = bits.Mul64(x60, 0xffffffffffffffff)
    var x78 uint64
    var x79 uint64
    x78, x79 = bits.Add64(x77, x74, uint64(0x0))
    var x80 uint64
    var x81 uint64
    x80, x81 = bits.Add64(x75, x72, uint64(p256Uint1(x79)))
    var x82 uint64
    var x83 uint64
    x82, x83 = bits.Add64(x73, x70, uint64(p256Uint1(x81)))
    x84 := (uint64(p256Uint1(x83)) + x71)
    var x86 uint64
    _, x86 = bits.Add64(x60, x76, uint64(0x0))
    var x87 uint64
    var x88 uint64
    x87, x88 = bits.Add64(x62, x78, uint64(p256Uint1(x86)))
    var x89 uint64
    var x90 uint64
    x89, x90 = bits.Add64(x64, x80, uint64(p256Uint1(x88)))
    var x91 uint64
    var x92 uint64
    x91, x92 = bits.Add64(x66, x82, uint64(p256Uint1(x90)))
    var x93 uint64
    var x94 uint64
    x93, x94 = bits.Add64(x68, x84, uint64(p256Uint1(x92)))
    x95 := (uint64(p256Uint1(x94)) + uint64(p256Uint1(x69)))
    var x96 uint64
    var x97 uint64
    x97, x96 = bits.Mul64(x2, arg2[3])
    var x98 uint64
    var x99 uint64
    x99, x98 = bits.Mul64(x2, arg2[2])
    var x100 uint64
    var x101 uint64
    x101, x100 = bits.Mul64(x2, arg2[1])
    var x102 uint64
    var x103 uint64
    x103, x102 = bits.Mul64(x2, arg2[0])
    var x104 uint64
    var x105 uint64
    x104, x105 = bits.Add64(x103, x100, uint64(0x0))
    var x106 uint64
    var x107 uint64
    x106, x107 = bits.Add64(x101, x98, uint64(p256Uint1(x105)))
    var x108 uint64
    var x109 uint64
    x108, x109 = bits.Add64(x99, x96, uint64(p256Uint1(x107)))
    x110 := (uint64(p256Uint1(x109)) + x97)
    var x111 uint64
    var x112 uint64
    x111, x112 = bits.Add64(x87, x102, uint64(0x0))
    var x113 uint64
    var x114 uint64
    x113, x114 = bits.Add64(x89, x104, uint64(p256Uint1(x112)))
    var x115 uint64
    var x116 uint64
    x115, x116 = bits.Add64(x91, x106, uint64(p256Uint1(x114)))
    var x117 uint64
    var x118 uint64
    x117, x118 = bits.Add64(x93, x108, uint64(p256Uint1(x116)))
    var x119 uint64
    var x120 uint64
    x119, x120 = bits.Add64(x95, x110, uint64(p256Uint1(x118)))
    var x121 uint64
    var x122 uint64
    x122, x121 = bits.Mul64(x111, 0xfffffffeffffffff)
    var x123 uint64
    var x124 uint64
    x124, x123 = bits.Mul64(x111, 0xffffffffffffffff)
    var x125 uint64
    var x126 uint64
    x126, x125 = bits.Mul64(x111, 0xffffffff00000000)
    var x127 uint64
    var x128 uint64
    x128, x127 = bits.Mul64(x111, 0xffffffffffffffff)
    var x129 uint64
    var x130 uint64
    x129, x130 = bits.Add64(x128, x125, uint64(0x0))
    var x131 uint64
    var x132 uint64
    x131, x132 = bits.Add64(x126, x123, uint64(p256Uint1(x130)))
    var x133 uint64
    var x134 uint64
    x133, x134 = bits.Add64(x124, x121, uint64(p256Uint1(x132)))
    x135 := (uint64(p256Uint1(x134)) + x122)
    var x137 uint64
    _, x137 = bits.Add64(x111, x127, uint64(0x0))
    var x138 uint64
    var x139 uint64
    x138, x139 = bits.Add64(x113, x129, uint64(p256Uint1(x137)))
    var x140 uint64
    var x141 uint64
    x140, x141 = bits.Add64(x115, x131, uint64(p256Uint1(x139)))
    var x142 uint64
    var x143 uint64
    x142, x143 = bits.Add64(x117, x133, uint64(p256Uint1(x141)))
    var x144 uint64
    var x145 uint64
    x144, x145 = bits.Add64(x119, x135, uint64(p256Uint1(x143)))
    x146 := (uint64(p256Uint1(x145)) + uint64(p256Uint1(x120)))
    var x147 uint64
    var x148 uint64
    x148, x147 = bits.Mul64(x3, arg2[3])
    var x149 uint64
    var x150 uint64
    x150, x149 = bits.Mul64(x3, arg2[2])
    var x151 uint64
    var x152 uint64
    x152, x151 = bits.Mul64(x3, arg2[1])
    var x153 uint64
    var x154 uint64
    x154, x153 = bits.Mul64(x3, arg2[0])
    var x155 uint64
    var x156 uint64
    x155, x156 = bits.Add64(x154, x151, uint64(0x0))
    var x157 uint64
    var x158 uint64
    x157, x158 = bits.Add64(x152, x149, uint64(p256Uint1(x156)))
    var x159 uint64
    var x160 uint64
    x159, x160 = bits.Add64(x150, x147, uint64(p256Uint1(x158)))
    x161 := (uint64(p256Uint1(x160)) + x148)
    var x162 uint64
    var x163 uint64
    x162, x163 = bits.Add64(x138, x153, uint64(0x0))
    var x164 uint64
    var x165 uint64
    x164, x165 = bits.Add64(x140, x155, uint64(p256Uint1(x163)))
    var x166 uint64
    var x167 uint64
    x166, x167 = bits.Add64(x142, x157, uint64(p256Uint1(x165)))
    var x168 uint64
    var x169 uint64
    x168, x169 = bits.Add64(x144, x159, uint64(p256Uint1(x167)))
    var x170 uint64
    var x171 uint64
    x170, x171 = bits.Add64(x146, x161, uint64(p256Uint1(x169)))
    var x172 uint64
    var x173 uint64
    x173, x172 = bits.Mul64(x162, 0xfffffffeffffffff)
    var x174 uint64
    var x175 uint64
    x175, x174 = bits.Mul64(x162, 0xffffffffffffffff)
    var x176 uint64
    var x177 uint64
    x177, x176 = bits.Mul64(x162, 0xffffffff00000000)
    var x178 uint64
    var x179 uint64
    x179, x178 = bits.Mul64(x162, 0xffffffffffffffff)
    var x180 uint64
    var x181 uint64
    x180, x181 = bits.Add64(x179, x176, uint64(0x0))
    var x182 uint64
    var x183 uint64
    x182, x183 = bits.Add64(x177, x174, uint64(p256Uint1(x181)))
    var x184 uint64
    var x185 uint64
    x184, x185 = bits.Add64(x175, x172, uint64(p256Uint1(x183)))
    x186 := (uint64(p256Uint1(x185)) + x173)
    var x188 uint64
    _, x188 = bits.Add64(x162, x178, uint64(0x0))
    var x189 uint64
    var x190 uint64
    x189, x190 = bits.Add64(x164, x180, uint64(p256Uint1(x188)))
    var x191 uint64
    var x192 uint64
    x191, x192 = bits.Add64(x166, x182, uint64(p256Uint1(x190)))
    var x193 uint64
    var x194 uint64
    x193, x194 = bits.Add64(x168, x184, uint64(p256Uint1(x192)))
    var x195 uint64
    var x196 uint64
    x195, x196 = bits.Add64(x170, x186, uint64(p256Uint1(x194)))
    x197 := (uint64(p256Uint1(x196)) + uint64(p256Uint1(x171)))
    var x198 uint64
    var x199 uint64
    x198, x199 = bits.Sub64(x189, 0xffffffffffffffff, uint64(0x0))
    var x200 uint64
    var x201 uint64
    x200, x201 = bits.Sub64(x191, 0xffffffff00000000, uint64(p256Uint1(x199)))
    var x202 uint64
    var x203 uint64
    x202, x203 = bits.Sub64(x193, 0xffffffffffffffff, uint64(p256Uint1(x201)))
    var x204 uint64
    var x205 uint64
    x204, x205 = bits.Sub64(x195, 0xfffffffeffffffff, uint64(p256Uint1(x203)))
    var x207 uint64
    _, x207 = bits.Sub64(x197, uint64(0x0), uint64(p256Uint1(x205)))
    var x208 uint64
    p256CmovznzU64(&x208, p256Uint1(x207), x198, x189)
    var x209 uint64
    p256CmovznzU64(&x209, p256Uint1(x207), x200, x191)
    var x210 uint64
    p256CmovznzU64(&x210, p256Uint1(x207), x202, x193)
    var x211 uint64
    p256CmovznzU64(&x211, p256Uint1(x207), x204, x195)
    out1[0] = x208
    out1[1] = x209
    out1[2] = x210
    out1[3] = x211
}

// p256Square squares a field element in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg1)) mod m
//   0 ≤ eval out1 < m
//
func p256Square(out1 *[4]uint64, arg1 *[4]uint64) {
    x1 := arg1[1]
    x2 := arg1[2]
    x3 := arg1[3]
    x4 := arg1[0]
    var x5 uint64
    var x6 uint64
    x6, x5 = bits.Mul64(x4, arg1[3])
    var x7 uint64
    var x8 uint64
    x8, x7 = bits.Mul64(x4, arg1[2])
    var x9 uint64
    var x10 uint64
    x10, x9 = bits.Mul64(x4, arg1[1])
    var x11 uint64
    var x12 uint64
    x12, x11 = bits.Mul64(x4, arg1[0])
    var x13 uint64
    var x14 uint64
    x13, x14 = bits.Add64(x12, x9, uint64(0x0))
    var x15 uint64
    var x16 uint64
    x15, x16 = bits.Add64(x10, x7, uint64(p256Uint1(x14)))
    var x17 uint64
    var x18 uint64
    x17, x18 = bits.Add64(x8, x5, uint64(p256Uint1(x16)))
    x19 := (uint64(p256Uint1(x18)) + x6)
    var x20 uint64
    var x21 uint64
    x21, x20 = bits.Mul64(x11, 0xfffffffeffffffff)
    var x22 uint64
    var x23 uint64
    x23, x22 = bits.Mul64(x11, 0xffffffffffffffff)
    var x24 uint64
    var x25 uint64
    x25, x24 = bits.Mul64(x11, 0xffffffff00000000)
    var x26 uint64
    var x27 uint64
    x27, x26 = bits.Mul64(x11, 0xffffffffffffffff)
    var x28 uint64
    var x29 uint64
    x28, x29 = bits.Add64(x27, x24, uint64(0x0))
    var x30 uint64
    var x31 uint64
    x30, x31 = bits.Add64(x25, x22, uint64(p256Uint1(x29)))
    var x32 uint64
    var x33 uint64
    x32, x33 = bits.Add64(x23, x20, uint64(p256Uint1(x31)))
    x34 := (uint64(p256Uint1(x33)) + x21)
    var x36 uint64
    _, x36 = bits.Add64(x11, x26, uint64(0x0))
    var x37 uint64
    var x38 uint64
    x37, x38 = bits.Add64(x13, x28, uint64(p256Uint1(x36)))
    var x39 uint64
    var x40 uint64
    x39, x40 = bits.Add64(x15, x30, uint64(p256Uint1(x38)))
    var x41 uint64
    var x42 uint64
    x41, x42 = bits.Add64(x17, x32, uint64(p256Uint1(x40)))
    var x43 uint64
    var x44 uint64
    x43, x44 = bits.Add64(x19, x34, uint64(p256Uint1(x42)))
    var x45 uint64
    var x46 uint64
    x46, x45 = bits.Mul64(x1, arg1[3])
    var x47 uint64
    var x48 uint64
    x48, x47 = bits.Mul64(x1, arg1[2])
    var x49 uint64
    var x50 uint64
    x50, x49 = bits.Mul64(x1, arg1[1])
    var x51 uint64
    var x52 uint64
    x52, x51 = bits.Mul64(x1, arg1[0])
    var x53 uint64
    var x54 uint64
    x53, x54 = bits.Add64(x52, x49, uint64(0x0))
    var x55 uint64
    var x56 uint64
    x55, x56 = bits.Add64(x50, x47, uint64(p256Uint1(x54)))
    var x57 uint64
    var x58 uint64
    x57, x58 = bits.Add64(x48, x45, uint64(p256Uint1(x56)))
    x59 := (uint64(p256Uint1(x58)) + x46)
    var x60 uint64
    var x61 uint64
    x60, x61 = bits.Add64(x37, x51, uint64(0x0))
    var x62 uint64
    var x63 uint64
    x62, x63 = bits.Add64(x39, x53, uint64(p256Uint1(x61)))
    var x64 uint64
    var x65 uint64
    x64, x65 = bits.Add64(x41, x55, uint64(p256Uint1(x63)))
    var x66 uint64
    var x67 uint64
    x66, x67 = bits.Add64(x43, x57, uint64(p256Uint1(x65)))
    var x68 uint64
    var x69 uint64
    x68, x69 = bits.Add64(uint64(p256Uint1(x44)), x59, uint64(p256Uint1(x67)))
    var x70 uint64
    var x71 uint64
    x71, x70 = bits.Mul64(x60, 0xfffffffeffffffff)
    var x72 uint64
    var x73 uint64
    x73, x72 = bits.Mul64(x60, 0xffffffffffffffff)
    var x74 uint64
    var x75 uint64
    x75, x74 = bits.Mul64(x60, 0xffffffff00000000)
    var x76 uint64
    var x77 uint64
    x77, x76 = bits.Mul64(x60, 0xffffffffffffffff)
    var x78 uint64
    var x79 uint64
    x78, x79 = bits.Add64(x77, x74, uint64(0x0))
    var x80 uint64
    var x81 uint64
    x80, x81 = bits.Add64(x75, x72, uint64(p256Uint1(x79)))
    var x82 uint64
    var x83 uint64
    x82, x83 = bits.Add64(x73, x70, uint64(p256Uint1(x81)))
    x84 := (uint64(p256Uint1(x83)) + x71)
    var x86 uint64
    _, x86 = bits.Add64(x60, x76, uint64(0x0))
    var x87 uint64
    var x88 uint64
    x87, x88 = bits.Add64(x62, x78, uint64(p256Uint1(x86)))
    var x89 uint64
    var x90 uint64
    x89, x90 = bits.Add64(x64, x80, uint64(p256Uint1(x88)))
    var x91 uint64
    var x92 uint64
    x91, x92 = bits.Add64(x66, x82, uint64(p256Uint1(x90)))
    var x93 uint64
    var x94 uint64
    x93, x94 = bits.Add64(x68, x84, uint64(p256Uint1(x92)))
    x95 := (uint64(p256Uint1(x94)) + uint64(p256Uint1(x69)))
    var x96 uint64
    var x97 uint64
    x97, x96 = bits.Mul64(x2, arg1[3])
    var x98 uint64
    var x99 uint64
    x99, x98 = bits.Mul64(x2, arg1[2])
    var x100 uint64
    var x101 uint64
    x101, x100 = bits.Mul64(x2, arg1[1])
    var x102 uint64
    var x103 uint64
    x103, x102 = bits.Mul64(x2, arg1[0])
    var x104 uint64
    var x105 uint64
    x104, x105 = bits.Add64(x103, x100, uint64(0x0))
    var x106 uint64
    var x107 uint64
    x106, x107 = bits.Add64(x101, x98, uint64(p256Uint1(x105)))
    var x108 uint64
    var x109 uint64
    x108, x109 = bits.Add64(x99, x96, uint64(p256Uint1(x107)))
    x110 := (uint64(p256Uint1(x109)) + x97)
    var x111 uint64
    var x112 uint64
    x111, x112 = bits.Add64(x87, x102, uint64(0x0))
    var x113 uint64
    var x114 uint64
    x113, x114 = bits.Add64(x89, x104, uint64(p256Uint1(x112)))
    var x115 uint64
    var x116 uint64
    x115, x116 = bits.Add64(x91, x106, uint64(p256Uint1(x114)))
    var x117 uint64
    var x118 uint64
    x117, x118 = bits.Add64(x93, x108, uint64(p256Uint1(x116)))
    var x119 uint64
    var x120 uint64
    x119, x120 = bits.Add64(x95, x110, uint64(p256Uint1(x118)))
    var x121 uint64
    var x122 uint64
    x122, x121 = bits.Mul64(x111, 0xfffffffeffffffff)
    var x123 uint64
    var x124 uint64
    x124, x123 = bits.Mul64(x111, 0xffffffffffffffff)
    var x125 uint64
    var x126 uint64
    x126, x125 = bits.Mul64(x111, 0xffffffff00000000)
    var x127 uint64
    var x128 uint64
    x128, x127 = bits.Mul64(x111, 0xffffffffffffffff)
    var x129 uint64
    var x130 uint64
    x129, x130 = bits.Add64(x128, x125, uint64(0x0))
    var x131 uint64
    var x132 uint64
    x131, x132 = bits.Add64(x126, x123, uint64(p256Uint1(x130)))
    var x133 uint64
    var x134 uint64
    x133, x134 = bits.Add64(x124, x121, uint64(p256Uint1(x132)))
    x135 := (uint64(p256Uint1(x134)) + x122)
    var x137 uint64
    _, x137 = bits.Add64(x111, x127, uint64(0x0))
    var x138 uint64
    var x139 uint64
    x138, x139 = bits.Add64(x113, x129, uint64(p256Uint1(x137)))
    var x140 uint64
    var x141 uint64
    x140, x141 = bits.Add64(x115, x131, uint64(p256Uint1(x139)))
    var x142 uint64
    var x143 uint64
    x142, x143 = bits.Add64(x117, x133, uint64(p256Uint1(x141)))
    var x144 uint64
    var x145 uint64
    x144, x145 = bits.Add64(x119, x135, uint64(p256Uint1(x143)))
    x146 := (uint64(p256Uint1(x145)) + uint64(p256Uint1(x120)))
    var x147 uint64
    var x148 uint64
    x148, x147 = bits.Mul64(x3, arg1[3])
    var x149 uint64
    var x150 uint64
    x150, x149 = bits.Mul64(x3, arg1[2])
    var x151 uint64
    var x152 uint64
    x152, x151 = bits.Mul64(x3, arg1[1])
    var x153 uint64
    var x154 uint64
    x154, x153 = bits.Mul64(x3, arg1[0])
    var x155 uint64
    var x156 uint64
    x155, x156 = bits.Add64(x154, x151, uint64(0x0))
    var x157 uint64
    var x158 uint64
    x157, x158 = bits.Add64(x152, x149, uint64(p256Uint1(x156)))
    var x159 uint64
    var x160 uint64
    x159, x160 = bits.Add64(x150, x147, uint64(p256Uint1(x158)))
    x161 := (uint64(p256Uint1(x160)) + x148)
    var x162 uint64
    var x163 uint64
    x162, x163 = bits.Add64(x138, x153, uint64(0x0))
    var x164 uint64
    var x165 uint64
    x164, x165 = bits.Add64(x140, x155, uint64(p256Uint1(x163)))
    var x166 uint64
    var x167 uint64
    x166, x167 = bits.Add64(x142, x157, uint64(p256Uint1(x165)))
    var x168 uint64
    var x169 uint64
    x168, x169 = bits.Add64(x144, x159, uint64(p256Uint1(x167)))
    var x170 uint64
    var x171 uint64
    x170, x171 = bits.Add64(x146, x161, uint64(p256Uint1(x169)))
    var x172 uint64
    var x173 uint64
    x173, x172 = bits.Mul64(x162, 0xfffffffeffffffff)
    var x174 uint64
    var x175 uint64
    x175, x174 = bits.Mul64(x162, 0xffffffffffffffff)
    var x176 uint64
    var x177 uint64
    x177, x176 = bits.Mul64(x162, 0xffffffff00000000)
    var x178 uint64
    var x179 uint64
    x179, x178 = bits.Mul64(x162, 0xffffffffffffffff)
    var x180 uint64
    var x181 uint64
    x180, x181 = bits.Add64(x179, x176, uint64(0x0))
    var x182 uint64
    var x183 uint64
    x182, x183 = bits.Add64(x177, x174, uint64(p256Uint1(x181)))
    var x184 uint64
    var x185 uint64
    x184, x185 = bits.Add64(x175, x172, uint64(p256Uint1(x183)))
    x186 := (uint64(p256Uint1(x185)) + x173)
    var x188 uint64
    _, x188 = bits.Add64(x162, x178, uint64(0x0))
    var x189 uint64
    var x190 uint64
    x189, x190 = bits.Add64(x164, x180, uint64(p256Uint1(x188)))
    var x191 uint64
    var x192 uint64
    x191, x192 = bits.Add64(x166, x182, uint64(p256Uint1(x190)))
    var x193 uint64
    var x194 uint64
    x193, x194 = bits.Add64(x168, x184, uint64(p256Uint1(x192)))
    var x195 uint64
    var x196 uint64
    x195, x196 = bits.Add64(x170, x186, uint64(p256Uint1(x194)))
    x197 := (uint64(p256Uint1(x196)) + uint64(p256Uint1(x171)))
    var x198 uint64
    var x199 uint64
    x198, x199 = bits.Sub64(x189, 0xffffffffffffffff, uint64(0x0))
    var x200 uint64
    var x201 uint64
    x200, x201 = bits.Sub64(x191, 0xffffffff00000000, uint64(p256Uint1(x199)))
    var x202 uint64
    var x203 uint64
    x202, x203 = bits.Sub64(x193, 0xffffffffffffffff, uint64(p256Uint1(x201)))
    var x204 uint64
    var x205 uint64
    x204, x205 = bits.Sub64(x195, 0xfffffffeffffffff, uint64(p256Uint1(x203)))
    var x207 uint64
    _, x207 = bits.Sub64(x197, uint64(0x0), uint64(p256Uint1(x205)))
    var x208 uint64
    p256CmovznzU64(&x208, p256Uint1(x207), x198, x189)
    var x209 uint64
    p256CmovznzU64(&x209, p256Uint1(x207), x200, x191)
    var x210 uint64
    p256CmovznzU64(&x210, p256Uint1(x207), x202, x193)
    var x211 uint64
    p256CmovznzU64(&x211, p256Uint1(x207), x204, x195)
    out1[0] = x208
    out1[1] = x209
    out1[2] = x210
    out1[3] = x211
}

// p256Add adds two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) + eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func p256Add(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
    var x1 uint64
    var x2 uint64
    x1, x2 = bits.Add64(arg1[0], arg2[0], uint64(0x0))
    var x3 uint64
    var x4 uint64
    x3, x4 = bits.Add64(arg1[1], arg2[1], uint64(p256Uint1(x2)))
    var x5 uint64
    var x6 uint64
    x5, x6 = bits.Add64(arg1[2], arg2[2], uint64(p256Uint1(x4)))
    var x7 uint64
    var x8 uint64
    x7, x8 = bits.Add64(arg1[3], arg2[3], uint64(p256Uint1(x6)))
    var x9 uint64
    var x10 uint64
    x9, x10 = bits.Sub64(x1, 0xffffffffffffffff, uint64(0x0))
    var x11 uint64
    var x12 uint64
    x11, x12 = bits.Sub64(x3, 0xffffffff00000000, uint64(p256Uint1(x10)))
    var x13 uint64
    var x14 uint64
    x13, x14 = bits.Sub64(x5, 0xffffffffffffffff, uint64(p256Uint1(x12)))
    var x15 uint64
    var x16 uint64
    x15, x16 = bits.Sub64(x7, 0xfffffffeffffffff, uint64(p256Uint1(x14)))
    var x18 uint64
    _, x18 = bits.Sub64(uint64(p256Uint1(x8)), uint64(0x0), uint64(p256Uint1(x16)))
    var x19 uint64
    p256CmovznzU64(&x19, p256Uint1(x18), x9, x1)
    var x20 uint64
    p256CmovznzU64(&x20, p256Uint1(x18), x11, x3)
    var x21 uint64
    p256CmovznzU64(&x21, p256Uint1(x18), x13, x5)
    var x22 uint64
    p256CmovznzU64(&x22, p256Uint1(x18), x15, x7)
    out1[0] = x19
    out1[1] = x20
    out1[2] = x21
    out1[3] = x22
}

// p256Sub subtracts two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) - eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func p256Sub(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
    var x1 uint64
    var x2 uint64
    x1, x2 = bits.Sub64(arg1[0], arg2[0], uint64(0x0))
    var x3 uint64
    var x4 uint64
    x3, x4 = bits.Sub64(arg1[1], arg2[1], uint64(p256Uint1(x2)))
    var x5 uint64
    var x6 uint64
    x5, x6 = bits.Sub64(arg1[2], arg2[2], uint64(p256Uint1(x4)))
    var x7 uint64
    var x8 uint64
    x7, x8 = bits.Sub64(arg1[3], arg2[3], uint64(p256Uint1(x6)))
    var x9 uint64
    p256CmovznzU64(&x9, p256Uint1(x8), uint64(0x0), 0xffffffffffffffff)
    var x10 uint64
    var x11 uint64
    x10, x11 = bits.Add64(x1, x9, uint64(0x0))
    var x12 uint64
    var x13 uint64
    x12, x13 = bits.Add64(x3, (x9 & 0xffffffff00000000), uint64(p256Uint1(x11)))
    var x14 uint64
    var x15 uint64
    x14, x15 = bits.Add64(x5, x9, uint64(p256Uint1(x13)))
    var x16 uint64
    x16, _ = bits.Add64(x7, (x9 & 0xfffffffeffffffff), uint64(p256Uint1(x15)))
    out1[0] = x10
    out1[1] = x12
    out1[2] = x14
    out1[3] = x16
}

// p256SetOne returns the field element one in the Montgomery domain.
//
// Postconditions:
//   eval (from_montgomery out1) mod m = 1 mod m
//   0 ≤ eval out1 < m
//
func p256SetOne(out1 *[4]uint64) {
    out1[0] = uint64(0x1)
    out1[1] = 0xffffffff
    out1[2] = uint64(0x0)
    out1[3] = 0x100000000
}

// p256FromMontgomery translates a field element out of the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval out1 mod m = (eval arg1 * ((2^64)⁻¹ mod m)^4) mod m
//   0 ≤ eval out1 < m
//
func p256FromMontgomery(out1 *p256NonMontgomeryDomainFieldElement, arg1 *[4]uint64) {
    x1 := arg1[0]
    var x2 uint64
    var x3 uint64
    x3, x2 = bits.Mul64(x1, 0xfffffffeffffffff)
    var x4 uint64
    var x5 uint64
    x5, x4 = bits.Mul64(x1, 0xffffffffffffffff)
    var x6 uint64
    var x7 uint64
    x7, x6 = bits.Mul64(x1, 0xffffffff00000000)
    var x8 uint64
    var x9 uint64
    x9, x8 = bits.Mul64(x1, 0xffffffffffffffff)
    var x10 uint64
    var x11 uint64
    x10, x11 = bits.Add64(x9, x6, uint64(0x0))
    var x12 uint64
    var x13 uint64
    x12, x13 = bits.Add64(x7, x4, uint64(p256Uint1(x11)))
    var x14 uint64
    var x15 uint64
    x14, x15 = bits.Add64(x5, x2, uint64(p256Uint1(x13)))
    var x17 uint64
    _, x17 = bits.Add64(x1, x8, uint64(0x0))
    var x18 uint64
    var x19 uint64
    x18, x19 = bits.Add64(uint64(0x0), x10, uint64(p256Uint1(x17)))
    var x20 uint64
    var x21 uint64
    x20, x21 = bits.Add64(uint64(0x0), x12, uint64(p256Uint1(x19)))
    var x22 uint64
    var x23 uint64
    x22, x23 = bits.Add64(uint64(0x0), x14, uint64(p256Uint1(x21)))
    var x24 uint64
    var x25 uint64
    x24, x25 = bits.Add64(x18, arg1[1], uint64(0x0))
    var x26 uint64
    var x27 uint64
    x26, x27 = bits.Add64(x20, uint64(0x0), uint64(p256Uint1(x25)))
    var x28 uint64
    var x29 uint64
    x28, x29 = bits.Add64(x22, uint64(0x0), uint64(p256Uint1(x27)))
    var x30 uint64
    var x31 uint64
    x31, x30 = bits.Mul64(x24, 0xfffffffeffffffff)
    var x32 uint64
    var x33 uint64
    x33, x32 = bits.Mul64(x24, 0xffffffffffffffff)
    var x34 uint64
    var x35 uint64
    x35, x34 = bits.Mul64(x24, 0xffffffff00000000)
    var x36 uint64
    var x37 uint64
    x37, x36 = bits.Mul64(x24, 0xffffffffffffffff)
    var x38 uint64
    var x39 uint64
    x38, x39 = bits.Add64(x37, x34, uint64(0x0))
    var x40 uint64
    var x41 uint64
    x40, x41 = bits.Add64(x35, x32, uint64(p256Uint1(x39)))
    var x42 uint64
    var x43 uint64
    x42, x43 = bits.Add64(x33, x30, uint64(p256Uint1(x41)))
    var x45 uint64
    _, x45 = bits.Add64(x24, x36, uint64(0x0))
    var x46 uint64
    var x47 uint64
    x46, x47 = bits.Add64(x26, x38, uint64(p256Uint1(x45)))
    var x48 uint64
    var x49 uint64
    x48, x49 = bits.Add64(x28, x40, uint64(p256Uint1(x47)))
    var x50 uint64
    var x51 uint64
    x50, x51 = bits.Add64((uint64(p256Uint1(x29)) + (uint64(p256Uint1(x23)) + (uint64(p256Uint1(x15)) + x3))), x42, uint64(p256Uint1(x49)))
    var x52 uint64
    var x53 uint64
    x52, x53 = bits.Add64(x46, arg1[2], uint64(0x0))
    var x54 uint64
    var x55 uint64
    x54, x55 = bits.Add64(x48, uint64(0x0), uint64(p256Uint1(x53)))
    var x56 uint64
    var x57 uint64
    x56, x57 = bits.Add64(x50, uint64(0x0), uint64(p256Uint1(x55)))
    var x58 uint64
    var x59 uint64
    x59, x58 = bits.Mul64(x52, 0xfffffffeffffffff)
    var x60 uint64
    var x61 uint64
    x61, x60 = bits.Mul64(x52, 0xffffffffffffffff)
    var x62 uint64
    var x63 uint64
    x63, x62 = bits.Mul64(x52, 0xffffffff00000000)
    var x64 uint64
    var x65 uint64
    x65, x64 = bits.Mul64(x52, 0xffffffffffffffff)
    var x66 uint64
    var x67 uint64
    x66, x67 = bits.Add64(x65, x62, uint64(0x0))
    var x68 uint64
    var x69 uint64
    x68, x69 = bits.Add64(x63, x60, uint64(p256Uint1(x67)))
    var x70 uint64
    var x71 uint64
    x70, x71 = bits.Add64(x61, x58, uint64(p256Uint1(x69)))
    var x73 uint64
    _, x73 = bits.Add64(x52, x64, uint64(0x0))
    var x74 uint64
    var x75 uint64
    x74, x75 = bits.Add64(x54, x66, uint64(p256Uint1(x73)))
    var x76 uint64
    var x77 uint64
    x76, x77 = bits.Add64(x56, x68, uint64(p256Uint1(x75)))
    var x78 uint64
    var x79 uint64
    x78, x79 = bits.Add64((uint64(p256Uint1(x57)) + (uint64(p256Uint1(x51)) + (uint64(p256Uint1(x43)) + x31))), x70, uint64(p256Uint1(x77)))
    var x80 uint64
    var x81 uint64
    x80, x81 = bits.Add64(x74, arg1[3], uint64(0x0))
    var x82 uint64
    var x83 uint64
    x82, x83 = bits.Add64(x76, uint64(0x0), uint64(p256Uint1(x81)))
    var x84 uint64
    var x85 uint64
    x84, x85 = bits.Add64(x78, uint64(0x0), uint64(p256Uint1(x83)))
    var x86 uint64
    var x87 uint64
    x87, x86 = bits.Mul64(x80, 0xfffffffeffffffff)
    var x88 uint64
    var x89 uint64
    x89, x88 = bits.Mul64(x80, 0xffffffffffffffff)
    var x90 uint64
    var x91 uint64
    x91, x90 = bits.Mul64(x80, 0xffffffff00000000)
    var x92 uint64
    var x93 uint64
    x93, x92 = bits.Mul64(x80, 0xffffffffffffffff)
    var x94 uint64
    var x95 uint64
    x94, x95 = bits.Add64(x93, x90, uint64(0x0))
    var x96 uint64
    var x97 uint64
    x96, x97 = bits.Add64(x91, x88, uint64(p256Uint1(x95)))
    var x98 uint64
    var x99 uint64
    x98, x99 = bits.Add64(x89, x86, uint64(p256Uint1(x97)))
    var x101 uint64
    _, x101 = bits.Add64(x80, x92, uint64(0x0))
    var x102 uint64
    var x103 uint64
    x102, x103 = bits.Add64(x82, x94, uint64(p256Uint1(x101)))
    var x104 uint64
    var x105 uint64
    x104, x105 = bits.Add64(x84, x96, uint64(p256Uint1(x103)))
    var x106 uint64
    var x107 uint64
    x106, x107 = bits.Add64((uint64(p256Uint1(x85)) + (uint64(p256Uint1(x79)) + (uint64(p256Uint1(x71)) + x59))), x98, uint64(p256Uint1(x105)))
    x108 := (uint64(p256Uint1(x107)) + (uint64(p256Uint1(x99)) + x87))
    var x109 uint64
    var x110 uint64
    x109, x110 = bits.Sub64(x102, 0xffffffffffffffff, uint64(0x0))
    var x111 uint64
    var x112 uint64
    x111, x112 = bits.Sub64(x104, 0xffffffff00000000, uint64(p256Uint1(x110)))
    var x113 uint64
    var x114 uint64
    x113, x114 = bits.Sub64(x106, 0xffffffffffffffff, uint64(p256Uint1(x112)))
    var x115 uint64
    var x116 uint64
    x115, x116 = bits.Sub64(x108, 0xfffffffeffffffff, uint64(p256Uint1(x114)))
    var x118 uint64
    _, x118 = bits.Sub64(uint64(0x0), uint64(0x0), uint64(p256Uint1(x116)))
    var x119 uint64
    p256CmovznzU64(&x119, p256Uint1(x118), x109, x102)
    var x120 uint64
    p256CmovznzU64(&x120, p256Uint1(x118), x111, x104)
    var x121 uint64
    p256CmovznzU64(&x121, p256Uint1(x118), x113, x106)
    var x122 uint64
    p256CmovznzU64(&x122, p256Uint1(x118), x115, x108)
    out1[0] = x119
    out1[1] = x120
    out1[2] = x121
    out1[3] = x122
}

// p256ToMontgomery translates a field element into the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = eval arg1 mod m
//   0 ≤ eval out1 < m
//
func p256ToMontgomery(out1 *[4]uint64, arg1 *p256NonMontgomeryDomainFieldElement) {
    x1 := arg1[1]
    x2 := arg1[2]
    x3 := arg1[3]
    x4 := arg1[0]
    var x5 uint64
    var x6 uint64
    x6, x5 = bits.Mul64(x4, 0x400000002)
    var x7 uint64
    var x8 uint64
    x8, x7 = bits.Mul64(x4, 0x100000001)
    var x9 uint64
    var x10 uint64
    x10, x9 = bits.Mul64(x4, 0x2ffffffff)
    var x11 uint64
    var x12 uint64
    x12, x11 = bits.Mul64(x4, 0x200000003)
    var x13 uint64
    var x14 uint64
    x13, x14 = bits.Add64(x12, x9, uint64(0x0))
    var x15 uint64
    var x16 uint64
    x15, x16 = bits.Add64(x10, x7, uint64(p256Uint1(x14)))
    var x17 uint64
    var x18 uint64
    x17, x18 = bits.Add64(x8, x5, uint64(p256Uint1(x16)))
    var x19 uint64
    var x20 uint64
    x20, x19 = bits.Mul64(x11, 0xfffffffeffffffff)
    var x21 uint64
    var x22 uint64
    x22, x21 = bits.Mul64(x11, 0xffffffffffffffff)
    var x23 uint64
    var x24 uint64
    x24, x23 = bits.Mul64(x11, 0xffffffff00000000)
    var x25 uint64
    var x26 uint64
    x26, x25 = bits.Mul64(x11, 0xffffffffffffffff)
    var x27 uint64
    var x28 uint64
    x27, x28 = bits.Add64(x26, x23, uint64(0x0))
    var x29 uint64
    var x30 uint64
    x29, x30 = bits.Add64(x24, x21, uint64(p256Uint1(x28)))
    var x31 uint64
    var x32 uint64
    x31, x32 = bits.Add64(x22, x19, uint64(p256Uint1(x30)))
    var x34 uint64
    _, x34 = bits.Add64(x11, x25, uint64(0x0))
    var x35 uint64
    var x36 uint64
    x35, x36 = bits.Add64(x13, x27, uint64(p256Uint1(x34)))
    var x37 uint64
    var x38 uint64
    x37, x38 = bits.Add64(x15, x29, uint64(p256Uint1(x36)))
    var x39 uint64
    var x40 uint64
    x39, x40 = bits.Add64(x17, x31, uint64(p256Uint1(x38)))
    var x41 uint64
    var x42 uint64
    x41, x42 = bits.Add64((uint64(p256Uint1(x18)) + x6), (uint64(p256Uint1(x32)) + x20), uint64(p256Uint1(x40)))
    var x43 uint64
    var x44 uint64
    x44, x43 = bits.Mul64(x1, 0x400000002)
    var x45 uint64
    var x46 uint64
    x46, x45 = bits.Mul64(x1, 0x100000001)
    var x47 uint64
    var x48 uint64
    x48, x47 = bits.Mul64(x1, 0x2ffffffff)
    var x49 uint64
    var x50 uint64
    x50, x49 = bits.Mul64(x1, 0x200000003)
    var x51 uint64
    var x52 uint64
    x51, x52 = bits.Add64(x50, x47, uint64(0x0))
    var x53 uint64
    var x54 uint64
    x53, x54 = bits.Add64(x48, x45, uint64(p256Uint1(x52)))
    var x55 uint64
    var x56 uint64
    x55, x56 = bits.Add64(x46, x43, uint64(p256Uint1(x54)))
    var x57 uint64
    var x58 uint64
    x57, x58 = bits.Add64(x35, x49, uint64(0x0))
    var x59 uint64
    var x60 uint64
    x59, x60 = bits.Add64(x37, x51, uint64(p256Uint1(x58)))
    var x61 uint64
    var x62 uint64
    x61, x62 = bits.Add64(x39, x53, uint64(p256Uint1(x60)))
    var x63 uint64
    var x64 uint64
    x63, x64 = bits.Add64(x41, x55, uint64(p256Uint1(x62)))
    var x65 uint64
    var x66 uint64
    x66, x65 = bits.Mul64(x57, 0xfffffffeffffffff)
    var x67 uint64
    var x68 uint64
    x68, x67 = bits.Mul64(x57, 0xffffffffffffffff)
    var x69 uint64
    var x70 uint64
    x70, x69 = bits.Mul64(x57, 0xffffffff00000000)
    var x71 uint64
    var x72 uint64
    x72, x71 = bits.Mul64(x57, 0xffffffffffffffff)
    var x73 uint64
    var x74 uint64
    x73, x74 = bits.Add64(x72, x69, uint64(0x0))
    var x75 uint64
    var x76 uint64
    x75, x76 = bits.Add64(x70, x67, uint64(p256Uint1(x74)))
    var x77 uint64
    var x78 uint64
    x77, x78 = bits.Add64(x68, x65, uint64(p256Uint1(x76)))
    var x80 uint64
    _, x80 = bits.Add64(x57, x71, uint64(0x0))
    var x81 uint64
    var x82 uint64
    x81, x82 = bits.Add64(x59, x73, uint64(p256Uint1(x80)))
    var x83 uint64
    var x84 uint64
    x83, x84 = bits.Add64(x61, x75, uint64(p256Uint1(x82)))
    var x85 uint64
    var x86 uint64
    x85, x86 = bits.Add64(x63, x77, uint64(p256Uint1(x84)))
    var x87 uint64
    var x88 uint64
    x87, x88 = bits.Add64(((uint64(p256Uint1(x64)) + uint64(p256Uint1(x42))) + (uint64(p256Uint1(x56)) + x44)), (uint64(p256Uint1(x78)) + x66), uint64(p256Uint1(x86)))
    var x89 uint64
    var x90 uint64
    x90, x89 = bits.Mul64(x2, 0x400000002)
    var x91 uint64
    var x92 uint64
    x92, x91 = bits.Mul64(x2, 0x100000001)
    var x93 uint64
    var x94 uint64
    x94, x93 = bits.Mul64(x2, 0x2ffffffff)
    var x95 uint64
    var x96 uint64
    x96, x95 = bits.Mul64(x2, 0x200000003)
    var x97 uint64
    var x98 uint64
    x97, x98 = bits.Add64(x96, x93, uint64(0x0))
    var x99 uint64
    var x100 uint64
    x99, x100 = bits.Add64(x94, x91, uint64(p256Uint1(x98)))
    var x101 uint64
    var x102 uint64
    x101, x102 = bits.Add64(x92, x89, uint64(p256Uint1(x100)))
    var x103 uint64
    var x104 uint64
    x103, x104 = bits.Add64(x81, x95, uint64(0x0))
    var x105 uint64
    var x106 uint64
    x105, x106 = bits.Add64(x83, x97, uint64(p256Uint1(x104)))
    var x107 uint64
    var x108 uint64
    x107, x108 = bits.Add64(x85, x99, uint64(p256Uint1(x106)))
    var x109 uint64
    var x110 uint64
    x109, x110 = bits.Add64(x87, x101, uint64(p256Uint1(x108)))
    var x111 uint64
    var x112 uint64
    x112, x111 = bits.Mul64(x103, 0xfffffffeffffffff)
    var x113 uint64
    var x114 uint64
    x114, x113 = bits.Mul64(x103, 0xffffffffffffffff)
    var x115 uint64
    var x116 uint64
    x116, x115 = bits.Mul64(x103, 0xffffffff00000000)
    var x117 uint64
    var x118 uint64
    x118, x117 = bits.Mul64(x103, 0xffffffffffffffff)
    var x119 uint64
    var x120 uint64
    x119, x120 = bits.Add64(x118, x115, uint64(0x0))
    var x121 uint64
    var x122 uint64
    x121, x122 = bits.Add64(x116, x113, uint64(p256Uint1(x120)))
    var x123 uint64
    var x124 uint64
    x123, x124 = bits.Add64(x114, x111, uint64(p256Uint1(x122)))
    var x126 uint64
    _, x126 = bits.Add64(x103, x117, uint64(0x0))
    var x127 uint64
    var x128 uint64
    x127, x128 = bits.Add64(x105, x119, uint64(p256Uint1(x126)))
    var x129 uint64
    var x130 uint64
    x129, x130 = bits.Add64(x107, x121, uint64(p256Uint1(x128)))
    var x131 uint64
    var x132 uint64
    x131, x132 = bits.Add64(x109, x123, uint64(p256Uint1(x130)))
    var x133 uint64
    var x134 uint64
    x133, x134 = bits.Add64(((uint64(p256Uint1(x110)) + uint64(p256Uint1(x88))) + (uint64(p256Uint1(x102)) + x90)), (uint64(p256Uint1(x124)) + x112), uint64(p256Uint1(x132)))
    var x135 uint64
    var x136 uint64
    x136, x135 = bits.Mul64(x3, 0x400000002)
    var x137 uint64
    var x138 uint64
    x138, x137 = bits.Mul64(x3, 0x100000001)
    var x139 uint64
    var x140 uint64
    x140, x139 = bits.Mul64(x3, 0x2ffffffff)
    var x141 uint64
    var x142 uint64
    x142, x141 = bits.Mul64(x3, 0x200000003)
    var x143 uint64
    var x144 uint64
    x143, x144 = bits.Add64(x142, x139, uint64(0x0))
    var x145 uint64
    var x146 uint64
    x145, x146 = bits.Add64(x140, x137, uint64(p256Uint1(x144)))
    var x147 uint64
    var x148 uint64
    x147, x148 = bits.Add64(x138, x135, uint64(p256Uint1(x146)))
    var x149 uint64
    var x150 uint64
    x149, x150 = bits.Add64(x127, x141, uint64(0x0))
    var x151 uint64
    var x152 uint64
    x151, x152 = bits.Add64(x129, x143, uint64(p256Uint1(x150)))
    var x153 uint64
    var x154 uint64
    x153, x154 = bits.Add64(x131, x145, uint64(p256Uint1(x152)))
    var x155 uint64
    var x156 uint64
    x155, x156 = bits.Add64(x133, x147, uint64(p256Uint1(x154)))
    var x157 uint64
    var x158 uint64
    x158, x157 = bits.Mul64(x149, 0xfffffffeffffffff)
    var x159 uint64
    var x160 uint64
    x160, x159 = bits.Mul64(x149, 0xffffffffffffffff)
    var x161 uint64
    var x162 uint64
    x162, x161 = bits.Mul64(x149, 0xffffffff00000000)
    var x163 uint64
    var x164 uint64
    x164, x163 = bits.Mul64(x149, 0xffffffffffffffff)
    var x165 uint64
    var x166 uint64
    x165, x166 = bits.Add64(x164, x161, uint64(0x0))
    var x167 uint64
    var x168 uint64
    x167, x168 = bits.Add64(x162, x159, uint64(p256Uint1(x166)))
    var x169 uint64
    var x170 uint64
    x169, x170 = bits.Add64(x160, x157, uint64(p256Uint1(x168)))
    var x172 uint64
    _, x172 = bits.Add64(x149, x163, uint64(0x0))
    var x173 uint64
    var x174 uint64
    x173, x174 = bits.Add64(x151, x165, uint64(p256Uint1(x172)))
    var x175 uint64
    var x176 uint64
    x175, x176 = bits.Add64(x153, x167, uint64(p256Uint1(x174)))
    var x177 uint64
    var x178 uint64
    x177, x178 = bits.Add64(x155, x169, uint64(p256Uint1(x176)))
    var x179 uint64
    var x180 uint64
    x179, x180 = bits.Add64(((uint64(p256Uint1(x156)) + uint64(p256Uint1(x134))) + (uint64(p256Uint1(x148)) + x136)), (uint64(p256Uint1(x170)) + x158), uint64(p256Uint1(x178)))
    var x181 uint64
    var x182 uint64
    x181, x182 = bits.Sub64(x173, 0xffffffffffffffff, uint64(0x0))
    var x183 uint64
    var x184 uint64
    x183, x184 = bits.Sub64(x175, 0xffffffff00000000, uint64(p256Uint1(x182)))
    var x185 uint64
    var x186 uint64
    x185, x186 = bits.Sub64(x177, 0xffffffffffffffff, uint64(p256Uint1(x184)))
    var x187 uint64
    var x188 uint64
    x187, x188 = bits.Sub64(x179, 0xfffffffeffffffff, uint64(p256Uint1(x186)))
    var x190 uint64
    _, x190 = bits.Sub64(uint64(p256Uint1(x180)), uint64(0x0), uint64(p256Uint1(x188)))
    var x191 uint64
    p256CmovznzU64(&x191, p256Uint1(x190), x181, x173)
    var x192 uint64
    p256CmovznzU64(&x192, p256Uint1(x190), x183, x175)
    var x193 uint64
    p256CmovznzU64(&x193, p256Uint1(x190), x185, x177)
    var x194 uint64
    p256CmovznzU64(&x194, p256Uint1(x190), x187, x179)
    out1[0] = x191
    out1[1] = x192
    out1[2] = x193
    out1[3] = x194
}

// p256Selectznz is a multi-limb conditional select.
//
// Postconditions:
//   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
//   arg3: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
func p256Selectznz(out1 *[4]uint64, arg1 p256Uint1, arg2 *[4]uint64, arg3 *[4]uint64) {
    var x1 uint64
    p256CmovznzU64(&x1, arg1, arg2[0], arg3[0])
    var x2 uint64
    p256CmovznzU64(&x2, arg1, arg2[1], arg3[1])
    var x3 uint64
    p256CmovznzU64(&x3, arg1, arg2[2], arg3[2])
    var x4 uint64
    p256CmovznzU64(&x4, arg1, arg2[3], arg3[3])
    out1[0] = x1
    out1[1] = x2
    out1[2] = x3
    out1[3] = x4
}

// p256ToBytes serializes a field element NOT in the Montgomery domain to bytes in little-endian order.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..31]
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff]]
func p256ToBytes(out1 *[32]uint8, arg1 *[4]uint64) {
    x1 := arg1[3]
    x2 := arg1[2]
    x3 := arg1[1]
    x4 := arg1[0]
    x5 := (uint8(x4) & 0xff)
    x6 := (x4 >> 8)
    x7 := (uint8(x6) & 0xff)
    x8 := (x6 >> 8)
    x9 := (uint8(x8) & 0xff)
    x10 := (x8 >> 8)
    x11 := (uint8(x10) & 0xff)
    x12 := (x10 >> 8)
    x13 := (uint8(x12) & 0xff)
    x14 := (x12 >> 8)
    x15 := (uint8(x14) & 0xff)
    x16 := (x14 >> 8)
    x17 := (uint8(x16) & 0xff)
    x18 := uint8((x16 >> 8))
    x19 := (uint8(x3) & 0xff)
    x20 := (x3 >> 8)
    x21 := (uint8(x20) & 0xff)
    x22 := (x20 >> 8)
    x23 := (uint8(x22) & 0xff)
    x24 := (x22 >> 8)
    x25 := (uint8(x24) & 0xff)
    x26 := (x24 >> 8)
    x27 := (uint8(x26) & 0xff)
    x28 := (x26 >> 8)
    x29 := (uint8(x28) & 0xff)
    x30 := (x28 >> 8)
    x31 := (uint8(x30) & 0xff)
    x32 := uint8((x30 >> 8))
    x33 := (uint8(x2) & 0xff)
    x34 := (x2 >> 8)
    x35 := (uint8(x34) & 0xff)
    x36 := (x34 >> 8)
    x37 := (uint8(x36) & 0xff)
    x38 := (x36 >> 8)
    x39 := (uint8(x38) & 0xff)
    x40 := (x38 >> 8)
    x41 := (uint8(x40) & 0xff)
    x42 := (x40 >> 8)
    x43 := (uint8(x42) & 0xff)
    x44 := (x42 >> 8)
    x45 := (uint8(x44) & 0xff)
    x46 := uint8((x44 >> 8))
    x47 := (uint8(x1) & 0xff)
    x48 := (x1 >> 8)
    x49 := (uint8(x48) & 0xff)
    x50 := (x48 >> 8)
    x51 := (uint8(x50) & 0xff)
    x52 := (x50 >> 8)
    x53 := (uint8(x52) & 0xff)
    x54 := (x52 >> 8)
    x55 := (uint8(x54) & 0xff)
    x56 := (x54 >> 8)
    x57 := (uint8(x56) & 0xff)
    x58 := (x56 >> 8)
    x59 := (uint8(x58) & 0xff)
    x60 := uint8((x58 >> 8))
    out1[0] = x5
    out1[1] = x7
    out1[2] = x9
    out1[3] = x11
    out1[4] = x13
    out1[5] = x15
    out1[6] = x17
    out1[7] = x18
    out1[8] = x19
    out1[9] = x21
    out1[10] = x23
    out1[11] = x25
    out1[12] = x27
    out1[13] = x29
    out1[14] = x31
    out1[15] = x32
    out1[16] = x33
    out1[17] = x35
    out1[18] = x37
    out1[19] = x39
    out1[20] = x41
    out1[21] = x43
    out1[22] = x45
    out1[23] = x46
    out1[24] = x47
    out1[25] = x49
    out1[26] = x51
    out1[27] = x53
    out1[28] = x55
    out1[29] = x57
    out1[30] = x59
    out1[31] = x60
}

// p256FromBytes deserializes a field element NOT in the Montgomery domain from bytes in little-endian order.
//
// Preconditions:
//   0 ≤ bytes_eval arg1 < m
// Postconditions:
//   eval out1 mod m = bytes_eval arg1 mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
func p256FromBytes(out1 *[4]uint64, arg1 *[32]uint8) {
    x1 := (uint64(arg1[31]) << 56)
    x2 := (uint64(arg1[30]) << 48)
    x3 := (uint64(arg1[29]) << 40)
    x4 := (uint64(arg1[28]) << 32)
    x5 := (uint64(arg1[27]) << 24)
    x6 := (uint64(arg1[26]) << 16)
    x7 := (uint64(arg1[25]) << 8)
    x8 := arg1[24]
    x9 := (uint64(arg1[23]) << 56)
    x10 := (uint64(arg1[22]) << 48)
    x11 := (uint64(arg1[21]) << 40)
    x12 := (uint64(arg1[20]) << 32)
    x13 := (uint64(arg1[19]) << 24)
    x14 := (uint64(arg1[18]) << 16)
    x15 := (uint64(arg1[17]) << 8)
    x16 := arg1[16]
    x17 := (uint64(arg1[15]) << 56)
    x18 := (uint64(arg1[14]) << 48)
    x19 := (uint64(arg1[13]) << 40)
    x20 := (uint64(arg1[12]) << 32)
    x21 := (uint64(arg1[11]) << 24)
    x22 := (uint64(arg1[10]) << 16)
    x23 := (uint64(arg1[9]) << 8)
    x24 := arg1[8]
    x25 := (uint64(arg1[7]) << 56)
    x26 := (uint64(arg1[6]) << 48)
    x27 := (uint64(arg1[5]) << 40)
    x28 := (uint64(arg1[4]) << 32)
    x29 := (uint64(arg1[3]) << 24)
    x30 := (uint64(arg1[2]) << 16)
    x31 := (uint64(arg1[1]) << 8)
    x32 := arg1[0]
    x33 := (x31 + uint64(x32))
    x34 := (x30 + x33)
    x35 := (x29 + x34)
    x36 := (x28 + x35)
    x37 := (x27 + x36)
    x38 := (x26 + x37)
    x39 := (x25 + x38)
    x40 := (x23 + uint64(x24))
    x41 := (x22 + x40)
    x42 := (x21 + x41)
    x43 := (x20 + x42)
    x44 := (x19 + x43)
    x45 := (x18 + x44)
    x46 := (x17 + x45)
    x47 := (x15 + uint64(x16))
    x48 := (x14 + x47)
    x49 := (x13 + x48)
    x50 := (x12 + x49)
    x51 := (x11 + x50)
    x52 := (x10 + x51)
    x53 := (x9 + x52)
    x54 := (x7 + uint64(x8))
    x55 := (x6 + x54)
    x56 := (x5 + x55)
    x57 := (x4 + x56)
    x58 := (x3 + x57)
    x59 := (x2 + x58)
    x60 := (x1 + x59)
    out1[0] = x39
    out1[1] = x46
    out1[2] = x53
    out1[3] = x60
}
