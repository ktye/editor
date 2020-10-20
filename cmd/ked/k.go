package main

import "math"

func ini(x0 I) (r I) {
	//defer func(){fmt.Printf("ini: r=%x\n", r)}()
	var x1, x2 I
	_, _ = x1, x2
	MJ[0>>3] = J(289360742959022340)
	MI[12>>2] = I(1887966018)
	MI[128>>2] = I(x0)
	x1 = I(256)
	x2 = I(8)
	for x2 < x0 {
		MI[(4*x2)>>2] = I(x1)
		x1 = I(x1 * 2)
		x2 = I(x2 + 1)
	}
	MI[132>>2] = I(enl(mk(1, 0)))
	MI[136>>2] = I(enl(0))
	MI[148>>2] = I(cat(cat(mks(120), mks(121)), mks(122)))
	return x0
}
func bk(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("bk: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(32 - clz32((7 + (x1 * I(MC[x0])))))
	if x2 < 4 {
		return 4
	}
	return x2
}
func mk(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mk: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(bk(x0, x1))
	x3 = I(4 * x2)
	x4 = I(4 * MI[128>>2])
	for 0 == MI[x3>>2] {
		if SI(x3) >= SI(x4) {
			x4 = I(grow((1 + (x3 / 4))))
			MI[128>>2] = I(x4)
			MI[x3>>2] = I((1 << (x3 >> 2)))
			x4 = I(x3)
			x3 = I(x3 - 4)
		}
		x3 = I(x3 + 4)
	}
	if 128 == x3 {
		panic("trap")
	}
	x5 = I(MI[x3>>2])
	MI[x3>>2] = I(MI[x5>>2])
	x6 = I(x3 - 4)
	for x6 >= (4 * x2) {
		x7 = I(x5 + (1 << (x6 >> 2)))
		MI[x7>>2] = I(MI[x6>>2])
		MI[x6>>2] = I(x7)
		x6 = I(x6 - 4)
	}
	MI[x5>>2] = I((x1 | (x0 << 29)))
	MI[(x5+4)>>2] = I(1)
	return x5
}
func mki(x0 I) (r I) {
	//defer func(){fmt.Printf("mki: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(mk(2, 1))
	MI[(x1+8)>>2] = I(x0)
	return x1
}
func mkf(x0 F) (r I) {
	//defer func(){fmt.Printf("mkf: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(mk(3, 1))
	MF[(x1+8)>>3] = F(x0)
	return x1
}
func mkd(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mkd: r=%x\n", r)}()
	var x2 I
	_ = x2
	if 5 != tp(x0) {
		x2 = I(mk(4, 1))
		MF[(x2+8)>>3] = F(F(0))
		MF[(x2+16)>>3] = F(F(1))
		return add(x0, mul(x1, x2))
	}
	x2 = I(nn(x0))
	if 1 == x2 {
		if 1 != nn(x1) {
			x1 = I(enl(x1))
		}
	}
	x1 = I(lx(x1))
	if x2 != nn(x1) {
		panic("trap")
	}
	x2 = I(l2(x0, x1))
	MI[x2>>2] = I((2 | (7 << 29)))
	return x2
}
func mkc(x0 I) (r I) {
	//defer func(){fmt.Printf("mkc: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(mk(1, 1))
	MC[(x1 + 8)] = C(C(x0))
	return x1
}
func mks(x0 I) (r I) {
	//defer func(){fmt.Printf("mks: r=%x\n", r)}()
	return sc(mkc(x0))
}
func mkz(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mkz: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(mkd(x0, x1))
	MI[x2>>2] = I((2 | (6 << 29)))
	return x2
}
func l2(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("l2: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(mk(6, 2))
	MI[(x2+8)>>2] = I(x0)
	MI[(x2+12)>>2] = I(x1)
	return x2
}
func l3(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("l3: r=%x\n", r)}()
	var x3 I
	_ = x3
	x3 = I(mk(6, 3))
	MI[(x3+8)>>2] = I(x0)
	MI[(x3+12)>>2] = I(x1)
	MI[(x3+16)>>2] = I(x2)
	return x3
}
func nn(x0 I) (r I) {
	//defer func(){fmt.Printf("nn: r=%x\n", r)}()
	if x0 < 256 {
		return 1
	}
	return (536870911 & MI[x0>>2])
}
func tp(x0 I) (r I) {
	//defer func(){fmt.Printf("tp: r=%x\n", r)}()
	if x0 < 256 {
		return 0
	}
	return (MI[x0>>2] >> 29)
}
func fr(x0 I) {
	//defer func(){fmt.Printf("fr: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	x4 = I(4 * bk(x1, x2))
	MI[x0>>2] = I(MI[x4>>2])
	MI[x4>>2] = I(x0)
}
func dx(x0 I) {
	//defer func(){fmt.Printf("dx: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	if x0 > 255 {
		x1 = I(MI[(x0+4)>>2])
		MI[(x0+4)>>2] = I((x1 - 1))
		if 1 == x1 {
			x2 = I(tp(x0))
			x3 = I(nn(x0))
			x4 = I(8 + x0)
			if 0 != (n32(x2) + i32b((x2 > 5))) {
				for x5 = 0; x5 < x3; x5++ {
					dx(MI[(x4+(4*x5))>>2])
				}
			}
			fr(x0)
		}
	}
}
func rx(x0 I) {
	//defer func(){fmt.Printf("rx: r=%x\n", r)}()
	rxn(x0, 1)
}
func rxn(x0 I, x1 I) {
	//defer func(){fmt.Printf("rxn: r=%x\n", r)}()
	if x0 > 255 {
		x0 = I(x0 + 4)
		MI[x0>>2] = I((x1 + MI[x0>>2]))
	}
}
func rl(x0 I) {
	//defer func(){fmt.Printf("rl: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	for x4 = 0; x4 < x2; x4++ {
		rx(MI[x3>>2])
		x3 = I(x3 + 4)
	}
}
func rld(x0 I) {
	//defer func(){fmt.Printf("rld: r=%x\n", r)}()
	rl(x0)
	dx(x0)
}
func lx(x0 I) (r I) {
	//defer func(){fmt.Printf("lx: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 == 6 {
		return x0
	}
	if 0 != (i32b((x1 == 7)) + i32b((x2 == 1))) {
		return enl(x0)
	}
	if 0 == x1 {
		panic("trap")
	}
	x4 = I(mk(6, x2))
	x5 = I(x4 + 8)
	rxn(x0, x2)
	for x6 = 0; x6 < x2; x6++ {
		MI[x5>>2] = I(atx(x0, mki(x6)))
		x5 = I(x5 + 4)
	}
	dx(x0)
	return x4
}
func til(x0 I) (r I) {
	//defer func(){fmt.Printf("til: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if 4 == x1 {
		return zim(x0)
	}
	if 6 == x1 {
		return ech(x0, 161)
	}
	if 7 == x1 {
		x4 = I(MI[x3>>2])
		rx(x4)
		dx(x0)
		return x4
	}
	if 2 != x1 {
		panic("trap")
	}
	x5 = I(MI[x3>>2])
	dx(x0)
	if SI(x5) < SI(0) {
		return tir(-(x5))
	}
	return seq(0, x5, 1)
}
func seq(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("seq: r=%x\n", r)}()
	var x3, x4, x5 I
	_, _, _ = x3, x4, x5
	x3 = I(mk(2, x1))
	x4 = I(8 + x3)
	for x5 = 0; x5 < x1; x5++ {
		MI[x4>>2] = I((x2 * (x5 + x0)))
		x4 = I(x4 + 4)
	}
	return x3
}
func tir(x0 I) (r I) {
	//defer func(){fmt.Printf("tir: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	x1 = I(mk(2, x0))
	x2 = I(4 + (x1 + (4 * x0)))
	for x3 = 0; x3 < x0; x3++ {
		MI[x2>>2] = I(x3)
		x2 = I(x2 - 4)
	}
	return x1
}
func upx(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("upx: r=%x\n", r)}()
	var x2, x3, x4 I
	_, _, _ = x2, x3, x4
	x2 = I(tp(x0))
	x3 = I(tp(x1))
	if x2 == x3 {
		return x0
	}
	if 0 != (i32b((x2 == 7)) + i32b((x3 == 7))) {
		panic("trap")
	}
	if x3 == 6 {
		return lx(x0)
	}
	x4 = I(nn(x0))
	for x2 < x3 {
		x0 = I(up(x0, x2, x4))
		x2 = I(x2 + 1)
	}
	return x0
}
func up(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("up: r=%x\n", r)}()
	var x3, x4, x5, x6 I
	_, _, _, _ = x3, x4, x5, x6
	x3 = I(mk((x1 + 1), x2))
	x4 = I(x0 + 8)
	x5 = I(x3 + 8)
	switch x1 {
	case 1:
		for x6 = 0; x6 < x2; x6++ {
			MI[x5>>2] = I(I(MC[(x4 + x6)]))
			x5 = I(x5 + 4)
		}
	case 2:
		for x6 = 0; x6 < x2; x6++ {
			MF[x5>>3] = F(F(SI(MI[x4>>2])))
			x5 = I(x5 + 8)
			x4 = I(x4 + 4)
		}
	case 3:
		for x6 = 0; x6 < x2; x6++ {
			MF[x5>>3] = F(MF[x4>>3])
			MF[(x5+8)>>3] = F(0.0)
			x4 = I(x4 + 8)
			x5 = I(x5 + 16)
		}
	default:
		panic("trap")
	}
	dx(x0)
	return x3
}
func atx(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("atx: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 I
	_, _, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if 0 == x2 {
		return cal(x0, enl(x1))
	}
	if x2 == 7 {
		return atd(x0, x1, x5)
	}
	if x5 > 5 {
		return ecr(x0, x1, 64)
	}
	if x5 == 3 {
		if x2 < 5 {
			return phi(x0, x1)
		}
	}
	if x5 != 2 {
		panic("trap")
	}
	x8 = I(mk(x2, x6))
	x9 = I(x8 + 8)
	x10 = I(I(MC[x2]))
	for x12 = 0; x12 < x6; x12++ {
		x11 = I(MI[x7>>2])
		if x3 <= x11 {
			panic("trap")
		}
		mv(x9, (x4 + (x10 * x11)), x10)
		x9 = I(x9 + x10)
		x7 = I(x7 + 4)
	}
	if x2 > 4 {
		rl(x8)
	}
	if x6 == 1 {
		if x2 == 6 {
			rx(MI[(x8+8)>>2])
			dx(x8)
			x8 = I(MI[(x8+8)>>2])
		}
	}
	dx(x0)
	dx(x1)
	return x8
}
func atm(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("atm: r=%x\n", r)}()
	var x2, x3, x4 I
	_, _, _ = x2, x3, x4
	if 0 == nn(x1) {
		dx(x0)
		return x1
	}
	rx(x1)
	x2 = I(fst(x1))
	x3 = I(drop(x1, 1))
	if 1 < nn(x3) {
		return atm(atx(x0, x2), x3)
	}
	x3 = I(fst(x3))
	x4 = I(nn(x2))
	if 0 != (n32(x2) + n32(i32b((x4 == 1)))) {
		if 0 != x2 {
			x0 = I(atx(x0, x2))
		}
		return ecl(x0, x3, 64)
	}
	return atx(atx(x0, x2), x3)
}
func atd(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("atd: r=%x\n", r)}()
	var x3, x4 I
	_, _ = x3, x4
	x3 = I(MI[(x0+8)>>2])
	x4 = I(MI[(x0+12)>>2])
	if x2 == 5 {
		rx(x3)
		x1 = I(fnd(x3, x1))
		x2 = I(2)
	}
	rx(x4)
	dx(x0)
	return atx(x4, x1)
}
func cal(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cal: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11 I
	_, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11
	x1 = I(lx(x1))
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if 0 != x2 {
		return atm(x0, x1)
	}
	if x6 == 1 {
		if 0 != (sadv(x0) | sadv((x0 - 128))) {
			return MT[x0].(func(I) I)(fst(x1))
		}
		if x0 < 128 {
			x0 = I(x0 + 128)
		}
	}
	if x0 < 128 {
		if x6 != 2 {
			panic("trap")
		}
		rld(x1)
		return MT[x0].(func(I, I) I)(MI[x7>>2], MI[(x7+4)>>2])
	}
	if x0 < 256 {
		if x6 != 1 {
			panic("trap")
		}
		return MT[x0].(func(I) I)(fst(x1))
	}
	if x3 == 2 {
		rld(x0)
		x8 = I(MI[x4>>2])
		switch x6 {
		case 1:
			return MT[(x8+128)].(func(I, I) I)(fst(x1), MI[(x4+4)>>2])
		case 2:
			rld(x1)
			return MT[x8].(func(I, I, I) I)(MI[x7>>2], MI[(x7+4)>>2], MI[(x4+4)>>2])
		default:
			panic("trap")
		}
	}
	if x3 == 3 {
		rl(x0)
		if 1 == x6 {
			x1 = I(fst(x1))
		}
		x9 = I(asi(MI[(x0+12)>>2], MI[(x0+16)>>2], x1))
		x10 = I(MI[(x0+8)>>2])
		dx(x0)
		return cal(x10, x9)
	}
	if x3 == 4 {
		x8 = I(MI[(x0+20)>>2])
		if x8 > x6 {
			x8 = I(x8 - x6)
			for x11 = 0; x11 < x8; x11++ {
				x1 = I(lcat(x1, 0))
			}
			return prj(x0, x1, seq(x6, x8, 1))
		}
		return lcl(x0, x1, 0)
	}
	panic("trap")
	return x0
}
func lcl(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("lcl: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14 I
	_, _, _, _, _, _, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14
	x3 = I(MI[(x0+20)>>2])
	if 0 == x3 {
		dx(x1)
		x1 = I(mk(6, 0))
	}
	if x3 != nn(x1) {
		panic("trap")
	}
	x4 = I(x1 + 8)
	x5 = I(MI[(x0+16)>>2])
	x6 = I(x5 + 8)
	x7 = I(nn(x5))
	x8 = I(mk(2, x7))
	x9 = I(x8 + 8)
	x10 = I(8 + MI[136>>2])
	for x14 = 0; x14 < x7; x14++ {
		x11 = I(x10 + (4 * MI[x6>>2]))
		MI[x9>>2] = I(MI[x11>>2])
		x12 = I(0)
		if x14 < x3 {
			x12 = I(MI[x4>>2])
			rx(x12)
			x4 = I(x4 + 4)
		}
		MI[x11>>2] = I(x12)
		x6 = I(x6 + 4)
		x9 = I(x9 + 4)
	}
	dx(x1)
	x13 = I(MI[(x0+12)>>2])
	rx(x13)
	x13 = I(evl(x13))
	x10 = I(8 + MI[136>>2])
	x9 = I(x8 + 8)
	x6 = I(x5 + 8)
	for x14 = 0; x14 < x7; x14++ {
		x11 = I(x10 + (4 * MI[x6>>2]))
		dx(MI[x11>>2])
		MI[x11>>2] = I(MI[x9>>2])
		x6 = I(x6 + 4)
		x9 = I(x9 + 4)
	}
	dx(x8)
	dx(x0)
	return x13
}
func rev(x0 I) (r I) {
	//defer func(){fmt.Printf("rev: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	if 7 == tp(x0) {
		rld(x0)
		x1 = I(MI[(x0+8)>>2])
		x2 = I(MI[(x0+12)>>2])
		return mkd(rev(x1), rev(x2))
	}
	x3 = I(nn(x0))
	if 0 == x3 {
		return x0
	}
	return atx(x0, tir(x3))
}
func fst(x0 I) (r I) {
	//defer func(){fmt.Printf("fst: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if 0 == x2 {
		dx(x0)
		if x1 == 0 {
			return 0
		}
		if x1 == 5 {
			return sc(mk(1, 0))
		}
		if x1 > 5 {
			return mk(6, 0)
		}
		return cst(mki(x1), mkc(0))
	}
	if 0 == x1 {
		return x0
	}
	if x1 == 7 {
		return fst(val(x0))
	}
	return atx(x0, mki(0))
}
func lst(x0 I) (r I) {
	//defer func(){fmt.Printf("lst: r=%x\n", r)}()
	if 7 == tp(x0) {
		return lst(val(x0))
	}
	return atx(x0, mki((nn(x0) - 1)))
}
func cut(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cut: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14 I
	_, _, _, _, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x5 == 7 {
		if x2 == 2 {
			if 1 != x3 {
				panic("trap")
			}
			rld(x1)
			x8 = I(MI[(x1+8)>>2])
			x9 = I(MI[(x1+12)>>2])
			rx(x0)
			return mkd(cut(x0, x8), cut(x0, x9))
		}
		rx(x1)
		return tkd(exc(til(x1), x0), x1)
	}
	if x2 != 2 {
		panic("trap")
	}
	if x3 == 1 {
		x10 = I(drop(x1, MI[x4>>2]))
		dx(x0)
		return x10
	}
	x10 = I(mk(6, x3))
	x11 = I(x10 + 8)
	for x14 = 0; x14 < x3; x14++ {
		x12 = I(MI[x4>>2])
		x13 = I(MI[(x4+4)>>2])
		if x14 == (x3 - 1) {
			x13 = I(x6)
		}
		if x13 < x12 {
			panic("trap")
		}
		rx(x1)
		MI[x11>>2] = I(atx(x1, seq(x12, (x13-x12), 1)))
		x4 = I(x4 + 4)
		x11 = I(x11 + 4)
	}
	dx(x0)
	dx(x1)
	return x10
}
func rsh(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("rsh: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14 I
	_, _, _, _, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x5 == 7 {
		if x2 == 2 {
			if 1 != x3 {
				panic("trap")
			}
			rld(x1)
			x8 = I(MI[(x1+8)>>2])
			x9 = I(MI[(x1+12)>>2])
			rx(x0)
			return mkd(rsh(x0, x8), rsh(x0, x9))
		}
		return tkd(x0, x1)
	}
	if x2 != 2 {
		panic("trap")
	}
	x10 = I(prod(x4, x3))
	x11 = I(take(x1, x10))
	if x3 == 1 {
		if x5 == 6 {
			if x10 == 1 {
				x11 = I(enl(x11))
			}
		}
		dx(x0)
		return x11
	}
	x3 = I(x3 - 1)
	x12 = I(x4 + (4 * x3))
	for x14 = 0; x14 < x3; x14++ {
		x13 = I(MI[x12>>2])
		x10 = I(x10 / x13)
		x10 = I(prod(x4, (x3 - x14)))
		x11 = I(cut(seq(0, x10, x13), x11))
		x12 = I(x12 - 4)
	}
	dx(x0)
	return x11
}
func prod(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("prod: r=%x\n", r)}()
	var x2, x3 I
	_, _ = x2, x3
	x2 = I(1)
	for x3 = 0; x3 < x1; x3++ {
		x2 = I(x2 * MI[x0>>2])
		x0 = I(x0 + 4)
	}
	return x2
}
func take(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("take: r=%x\n", r)}()
	var x2, x3, x4, x5, x6 I
	_, _, _, _, _ = x2, x3, x4, x5, x6
	x2 = I(nn(x0))
	x3 = I(0)
	if SI(x1) < SI(0) {
		x3 = I(x2 + x1)
		x1 = I(-(x1))
		if SI(x3) < SI(0) {
			return x0
		}
	}
	x4 = I(seq(x3, x1, 1))
	if x2 < x1 {
		x5 = I(x4 + 8)
		for x6 = 0; x6 < x1; x6++ {
			MI[x5>>2] = I((x6 % x2))
			x5 = I(x5 + 4)
		}
	}
	return atx(x0, x4)
}
func drop(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("drop: r=%x\n", r)}()
	var x2, x3, x4, x5 I
	_, _, _, _ = x2, x3, x4, x5
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(x1)
	if SI(x1) < SI(0) {
		x1 = I(0 - x1)
		x5 = I(0)
	}
	if x1 > x3 {
		dx(x0)
		return mk(x2, 0)
	}
	x0 = I(atx(x0, seq(x5, (x3-x1), 1)))
	if x2 == 6 {
		if 1 == (x3 - x1) {
			x0 = I(enl(x0))
		}
	}
	return x0
}
func tkd(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("tkd: r=%x\n", r)}()
	var x2, x3, x4 I
	_, _, _ = x2, x3, x4
	x2 = I(tp(x0))
	rld(x1)
	x3 = I(MI[(x1+8)>>2])
	x4 = I(MI[(x1+12)>>2])
	if x2 != 5 {
		panic("trap")
	}
	rx(x3)
	x0 = I(fnd(x3, x0))
	rx(x0)
	x4 = I(atx(x4, x0))
	if 1 == nn(x0) {
		x4 = I(enl(x4))
	}
	return mkd(atx(x3, x0), x4)
}
func phi(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("phi: r=%x\n", r)}()
	var x2, x3, x4, x5, x7 I

	var x6 F
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(nn(x1))
	x3 = I(mk(4, x2))
	x4 = I(x3 + 8)
	x5 = I(x1 + 8)
	for x7 = 0; x7 < x2; x7++ {
		x6 = F(0.017453292519943295 * MF[x5>>3])
		MF[x4>>3] = F(cos(x6))
		MF[(x4+8)>>3] = F(sin(x6))
		x4 = I(x4 + 16)
		x5 = I(x5 + 8)
	}
	dx(x1)
	return mul(x0, x3)
}
func use(x0 I) (r I) {
	//defer func(){fmt.Printf("use: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	if 1 == MI[(x0+4)>>2] {
		return x0
	}
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	x4 = I(mk(x1, x2))
	x5 = I(x4 + 8)
	mv(x5, x3, (x2 * I(MC[x1])))
	dx(x0)
	return x4
}
func mv(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("mv: r=%x\n", r)}()
	var x3 I
	_ = x3
	for x3 = 0; x3 < x2; x3++ {
		MC[(x0 + x3)] = C(C(I(MC[(x1 + x3)])))
	}
}
func cat(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cat: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if 0 == x2 {
		x0 = I(enl(x0))
		x2 = I(6)
	}
	if x2 == x5 {
		return ucat(x0, x1)
	}
	if x2 == 6 {
		return ucat(x0, lx(x1))
	}
	if x5 == 6 {
		return ucat(lx(x0), x1)
	}
	return ucat(lx(x0), lx(x1))
}
func ucat(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ucat: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x2 > 4 {
		rl(x0)
		rl(x1)
	}
	if x2 == 7 {
		x8 = I(mkd(ucat(MI[(x0+8)>>2], MI[(x1+8)>>2]), ucat(MI[(x0+12)>>2], MI[(x1+12)>>2])))
		dx(x0)
		dx(x1)
		return x8
	}
	x8 = I(mk(x2, (x3 + x6)))
	x9 = I(I(MC[x2]))
	mv((x8 + 8), x4, (x9 * x3))
	mv((x8 + (8 + (x9 * x3))), x7, (x9 * x6))
	dx(x0)
	dx(x1)
	return x8
}
func lcat(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("lcat: r=%x\n", r)}()
	var x2, x3, x4, x5 I
	_, _, _, _ = x2, x3, x4, x5
	x0 = I(use(x0))
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	if bk(x2, x3) < bk(x2, (x3+1)) {
		x5 = I(mk(x2, (x3 + 1)))
		rld(x0)
		mv((x5 + 8), x4, (4 * x3))
		x0 = I(x5)
		x4 = I(x0 + 8)
	}
	MI[(x4+(4*x3))>>2] = I(x1)
	MI[x0>>2] = I(((x3 + 1) | (6 << 29)))
	return x0
}
func enl(x0 I) (r I) {
	//defer func(){fmt.Printf("enl: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(mk(6, 1))
	MI[(x1+8)>>2] = I(x0)
	return x1
}
func cnt(x0 I) (r I) {
	//defer func(){fmt.Printf("cnt: r=%x\n", r)}()
	var x1 I
	_ = x1
	if 7 == tp(x0) {
		x0 = I(til(x0))
	}
	x1 = I(mki(nn(x0)))
	dx(x0)
	return x1
}
func typ(x0 I) (r I) {
	//defer func(){fmt.Printf("typ: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	x4 = I(mk(2, 1))
	MI[(8+x4)>>2] = I(x1)
	dx(x0)
	return x4
}
func not(x0 I) (r I) {
	//defer func(){fmt.Printf("not: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(tp(x0))
	if x1 > 5 {
		return ech(x0, 126)
	}
	if 0 == x1 {
		if 0 == x0 {
			return mki(1)
		}
		dx(x0)
		return mki(0)
	}
	return eql(x0, mki(0))
}
func wer(x0 I) (r I) {
	//defer func(){fmt.Printf("wer: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _, _, _ = x1, x2, x3, x4, x5, x6, x7, x8, x9
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 == 1 {
		return prs(x0)
	}
	if x1 == 4 {
		return zan(x0, x2, x3)
	}
	if x1 == 6 {
		return flp(x0)
	}
	if x1 != 2 {
		panic("trap")
	}
	x4 = I(0)
	for x7 = 0; x7 < x2; x7++ {
		x4 = I(x4 + MI[x3>>2])
		x3 = I(x3 + 4)
	}
	x3 = I(8 + x0)
	x5 = I(mk(2, x4))
	x6 = I(x5 + 8)
	for x7 = 0; x7 < x2; x7++ {
		x8 = MI[x3>>2]
		for x9 = 0; x9 < x8; x9++ {
			MI[x6>>2] = I(x7)
			x6 = I(x6 + 4)
		}
		x3 = I(x3 + 4)
	}
	dx(x0)
	return x5
}
func mtc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mtc: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(mk(2, 1))
	MI[(x2+8)>>2] = I(match(x0, x1))
	dx(x0)
	dx(x1)
	return x2
}
func match(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("match: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8 I
	_, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8
	if x0 == x1 {
		return 1
	}
	if MI[x0>>2] != MI[x1>>2] {
		return 0
	}
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(x1 + 8)
	x6 = I(0)
	switch x2 {
	case 0:
		return 1
	case 1:
		x7 = I(x3)
	case 2:
		x7 = I(x3 << 2)
	case 3:
		x7 = I(x3 << 3)
	case 4:
		x7 = I(x3 << 4)
	case 5:
		x7 = I(x3 << 2)
	default:
		for x8 = 0; x8 < x3; x8++ {
			if 0 == match(MI[x4>>2], MI[x5>>2]) {
				return 0
			}
			x4 = I(x4 + 4)
			x5 = I(x5 + 4)
		}
		return 1
	}
	for x8 = 0; x8 < x7; x8++ {
		if I(MC[(x4+x8)]) != I(MC[(x5+x8)]) {
			return 0
		}
	}
	return 1
}
func fnd(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("fnd: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11 I
	_, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x2 != x5 {
		panic("trap")
	}
	x8 = I(mk(2, x6))
	x9 = I(x8 + 8)
	x10 = I(I(MC[x5]))
	for x11 = 0; x11 < x6; x11++ {
		MI[x9>>2] = I(fnx(x0, x7))
		x9 = I(x9 + 4)
		x7 = I(x7 + x10)
	}
	dx(x0)
	dx(x1)
	return x8
}
func fnx(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("fnx: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(8 + x2)
	x6 = I(I(MC[x2]))
	for x7 = 0; x7 < x3; x7++ {
		if 0 != MT[x5].(func(I, I) I)(x4, x1) {
			return x7
		}
		x4 = I(x4 + x6)
	}
	return x3
}
func lop(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("lop: r=%x\n", r)}()
	var x3, x4 I
	_, _ = x3, x4
	x3 = I(tp(x1))
	if 0 == x3 {
		return fxp(x0, x1, x2)
	}
	if 6 == x3 {
		rld(x1)
		x4 = I(MI[(x1+12)>>2])
		x1 = I(MI[(x1+8)>>2])
		return whl(x1, x0, x4, x2)
	}
	dx(x2)
	return 0
}
func jon(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("jon: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(lop(x0, x1, 0))
	if 0 != x2 {
		return x2
	}
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	if 0 != (n32(i32b((x3 == 6))) + n32(x4)) {
		dx(x1)
		return x0
	}
	rl(x0)
	x2 = I(MI[x5>>2])
	rxn(x1, (x4 - 2))
	x6 = (x4 - 1)
	for x7 = 0; x7 < x6; x7++ {
		x5 = I(x5 + 4)
		x2 = I(cat(cat(x2, x1), MI[x5>>2]))
	}
	dx(x0)
	return x2
}
func spl(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("spl: r=%x\n", r)}()
	var x2, x3, x4, x5, x6 I
	_, _, _, _, _ = x2, x3, x4, x5, x6
	x2 = I(lop(x0, x1, enl(mk(6, 0))))
	if 0 != x2 {
		return x2
	}
	rx(x0)
	x3 = I(nn(x1))
	x2 = I(fds(x0, x1))
	if 0 == nn(x2) {
		dx(x2)
		return enl(x0)
	}
	x2 = I(cut(cat(mki(0), x2), x0))
	x4 = I(nn(x2) - 1)
	x5 = I(x2 + 8)
	for x6 = 0; x6 < x4; x6++ {
		x5 = I(x5 + 4)
		MI[x5>>2] = I(drop(MI[x5>>2], x3))
	}
	return x2
}
func fds(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("fds: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14 I
	_, _, _, _, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if 0 != (n32(i32b((x2 == x5))) + i32b((x2 > 5))) {
		panic("trap")
	}
	if x3 < x6 {
		dx(x0)
		dx(x1)
		return mk(2, 0)
	}
	if 0 == x6 {
		dx(x0)
		dx(x1)
		return drop(seq(0, x3, 1), 1)
	}
	x8 = I(mk(2, 0))
	x9 = I(I(MC[x2]))
	x10 = I(8 + x2)
	x11 = I(0)
	for x11 = 0; x11 < x3; x11++ {
		x12 = I(0)
		for x14 = 0; x14 < x6; x14++ {
			x13 = I(x9 * x14)
			x12 = I(x12 + MT[x10].(func(I, I) I)((x4+x13), (x7+x13)))
		}
		if x12 == x6 {
			x8 = I(ucat(x8, mki(x11)))
			x11 = I(x11 + (x6 - 1))
			x4 = I(x4 + (x9 * (x6 - 1)))
		}
		x4 = I(x4 + x9)
	}
	dx(x0)
	dx(x1)
	return x8
}
func exc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("exc: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(mki(nn(x1)))
	rx(x0)
	return atx(x0, wer(eql(x2, fnd(x1, x0))))
}
func srt(x0 I) (r I) {
	//defer func(){fmt.Printf("srt: r=%x\n", r)}()
	rx(x0)
	return atx(x0, grd(x0))
}
func gdn(x0 I) (r I) {
	//defer func(){fmt.Printf("gdn: r=%x\n", r)}()
	return rev(grd(x0))
}
func grd(x0 I) (r I) {
	//defer func(){fmt.Printf("grd: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	x4 = I(seq(0, x2, 1))
	x5 = I(seq(0, x2, 1))
	x6 = I(x4 + 8)
	msrt((x5 + 8), x6, 0, x2, x3, x1)
	dx(x0)
	dx(x5)
	return x4
}
func msrt(x0 I, x1 I, x2 I, x3 I, x4 I, x5 I) {
	//defer func(){fmt.Printf("msrt: r=%x\n", r)}()
	var x6 I
	_ = x6
	if (x3 - x2) >= 2 {
		x6 = I((x3 + x2) / 2)
		msrt(x1, x0, x2, x6, x4, x5)
		msrt(x1, x0, x6, x3, x4, x5)
		mrge(x0, x1, x2, x3, x6, x4, x5)
	}
}
func mrge(x0 I, x1 I, x2 I, x3 I, x4 I, x5 I, x6 I) {
	//defer func(){fmt.Printf("mrge: r=%x\n", r)}()
	var x7, x8, x9, x10, x11, x12 I
	_, _, _, _, _, _ = x7, x8, x9, x10, x11, x12
	x7 = I(x2)
	x8 = I(x4)
	x9 = I(I(MC[x6]))
	x10 = I(x2)
	for x10 < x3 {
		x11 = I(i32b((x7 >= x4)))
		if 0 == x11 {
			if x8 >= x3 {
				x11 = I(0)
			} else {
				x11 = I(MT[x6].(func(I, I) I)((x5 + (x9 * MI[(x0+(x7<<2))>>2])), (x5 + (x9 * MI[(x0+(x8<<2))>>2]))))
			}
		}
		if 0 != x11 {
			x12 = I(x8)
			x8 = I(x8 + 1)
		} else {
			x12 = I(x7)
			x7 = I(x7 + 1)
		}
		MI[(x1+(x10<<2))>>2] = I(MI[(x0+(x12<<2))>>2])
		x10 = I(x10 + 1)
	}
}
func gtc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("gtc: r=%x\n", r)}()
	return i32b((I(MC[x0]) > I(MC[x1])))
}
func gti(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("gti: r=%x\n", r)}()
	return i32b((SI(MI[x0>>2]) > SI(MI[x1>>2])))
}
func gtf(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("gtf: r=%x\n", r)}()
	return i32b((MF[x0>>3] > MF[x1>>3]))
}
func eqc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eqc: r=%x\n", r)}()
	return i32b((I(MC[x0]) == I(MC[x1])))
}
func eqi(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eqi: r=%x\n", r)}()
	return i32b((MI[x0>>2] == MI[x1>>2]))
}
func eqf(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eqf: r=%x\n", r)}()
	return i32b((MJ[x0>>3] == MJ[x1>>3]))
}
func eqz(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eqz: r=%x\n", r)}()
	if 0 != eqf(x0, x1) {
		return eqf((x0 + 8), (x1 + 8))
	}
	return 0
}
func eqL(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eqL: r=%x\n", r)}()
	return match(MI[x0>>2], MI[x1>>2])
}
func gtl(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("gtl: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 I
	_, _, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12
	x0 = I(MI[x0>>2])
	x1 = I(MI[x1>>2])
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x2 != x5 {
		return i32b((x2 > x5))
	}
	x8 = I(x3)
	if x6 < x3 {
		x8 = I(x6)
	}
	x9 = I(I(MC[x2]))
	for x12 = 0; x12 < x8; x12++ {
		x10 = I(x4 + (x12 * x9))
		x11 = I(x7 + (x12 * x9))
		if 0 != MT[x2].(func(I, I) I)(x10, x11) {
			return 1
		}
		if 0 != MT[x2].(func(I, I) I)(x11, x10) {
			return 0
		}
	}
	return i32b((x3 > x6))
}
func sc(x0 I) (r I) {
	//defer func(){fmt.Printf("sc: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	x1 = I(MI[132>>2])
	x2 = I(nn(x1))
	x0 = I(enl(x0))
	x3 = I(fnx(x1, (x0 + 8)))
	if x3 < x2 {
		dx(x0)
	} else {
		MI[132>>2] = I(cat(x1, x0))
		MI[136>>2] = I(lcat(MI[136>>2], 0))
	}
	x3 = I(mki(x3))
	MI[x3>>2] = I((1 | (5 << 29)))
	return x3
}
func cs(x0 I) (r I) {
	//defer func(){fmt.Printf("cs: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(MI[(x0+8)>>2])
	x1 = I(MI[(8+(MI[132>>2]+(4*x1)))>>2])
	rx(x1)
	dx(x0)
	return x1
}
func eql(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("eql: r=%x\n", r)}()
	return cmp(x0, x1, 1)
}
func mor(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mor: r=%x\n", r)}()
	return cmp(x0, x1, 0)
}
func les(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("les: r=%x\n", r)}()
	return cmp(x1, x0, 0)
}
func cmp(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("cmp: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13 I
	_, _, _, _, _, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13
	x0 = I(upx(x0, x1))
	x1 = I(upx(x1, x0))
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	x6 = I(tp(x1))
	x7 = I(nn(x1))
	x8 = I(8 + x1)
	if x4 != x7 {
		if x4 == 1 {
			x0 = I(take(x0, x7))
			x4 = I(x7)
			x5 = I(x0 + 8)
		}
		if x7 == 1 {
			x1 = I(take(x1, x4))
			x7 = I(x4)
			x8 = I(x1 + 8)
		}
	}
	if x3 == 6 {
		return ecd(x0, x1, (62 - x2))
	}
	x9 = I(x3)
	if 0 != x2 {
		x9 = I(x9 + 8)
	}
	x10 = I(I(MC[x3]))
	x11 = I(mk(2, x4))
	x12 = I(x11 + 8)
	for x13 = 0; x13 < x4; x13++ {
		MI[x12>>2] = I(MT[x9].(func(I, I) I)(x5, x8))
		x5 = I(x5 + x10)
		x8 = I(x8 + x10)
		x12 = I(x12 + 4)
	}
	dx(x0)
	dx(x1)
	return x11
}
func min(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("min: r=%x\n", r)}()
	return mia(x0, x1, 38)
}
func max(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("max: r=%x\n", r)}()
	return mia(x0, x1, 124)
}
func mia(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("mia: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9
	x0 = I(upx(x0, x1))
	x1 = I(upx(x1, x0))
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	x6 = I(tp(x1))
	x7 = I(nn(x1))
	x8 = I(8 + x1)
	if x4 != x7 {
		if x4 == 1 {
			x0 = I(take(x0, x7))
			x4 = I(x7)
			x5 = I(x0 + 8)
		}
		if x7 == 1 {
			x1 = I(take(x1, x4))
			x7 = I(x4)
			x8 = I(x1 + 8)
		}
	}
	if x3 == 6 {
		return ecd(x0, x1, x2)
	}
	rx(x0)
	rx(x1)
	if x2 == 38 {
		x9 = I(les(x0, x1))
	} else {
		x9 = I(mor(x0, x1))
	}
	x9 = I(wer(x9))
	rx(x9)
	return asi(x1, x9, atx(x0, x9))
}
func nd(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("nd: r=%x\n", r)}()
	var x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14 I
	_, _, _, _, _, _, _, _, _, _, _ = x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14
	x0 = I(upx(x0, x1))
	x1 = I(upx(x1, x0))
	x4 = I(tp(x0))
	x5 = I(nn(x0))
	x6 = I(8 + x0)
	x7 = I(tp(x1))
	x8 = I(nn(x1))
	x9 = I(8 + x1)
	if x5 != x8 {
		if x5 == 1 {
			x0 = I(take(x0, x8))
			x5 = I(x8)
			x6 = I(x0 + 8)
		}
		if x8 == 1 {
			x1 = I(take(x1, x5))
			x8 = I(x5)
			x9 = I(x1 + 8)
		}
	}
	if x4 == 6 {
		return ecd(x0, x1, x3)
	}
	x10 = I(I(MC[x4]))
	x11 = I(x2 + x4)
	x12 = I(mk(x4, x5))
	x13 = I(x12 + 8)
	for x14 = 0; x14 < x5; x14++ {
		MT[x11].(func(I, I, I))(x6, x9, x13)
		x6 = I(x6 + x10)
		x9 = I(x9 + x10)
		x13 = I(x13 + x10)
	}
	dx(x0)
	dx(x1)
	return x12
}
func nm(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("nm: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	if x3 > 5 {
		return ech(x0, x2)
	}
	x6 = I(use(x0))
	x7 = I(x6 + 8)
	x8 = I(I(MC[x3]))
	x1 = I(x1 + x3)
	for x9 = 0; x9 < x4; x9++ {
		MT[x1].(func(I, I))(x5, x7)
		x5 = I(x5 + x8)
		x7 = I(x7 + x8)
	}
	if x3 == 4 {
		if x1 == 19 {
			return zre(x6)
		}
	}
	return x6
}
func nmf(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("nmf: r=%x\n", r)}()
	var x2, x3, x4 I
	_, _, _ = x2, x3, x4
	dx(x0)
	x0 = I(MI[(x0+8)>>2])
	x1 = I(use(cst(mki(3), x1)))
	x2 = I(x1 + 8)
	x3 = nn(x1)
	for x4 = 0; x4 < x3; x4++ {
		MF[x2>>3] = F(MT[x0].(func(F) F)(MF[x2>>3]))
		x2 = I(x2 + 8)
	}
	return x1
}
func add(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("add: r=%x\n", r)}()
	return nd(x0, x1, 143, 43)
}
func sub(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("sub: r=%x\n", r)}()
	return nd(x0, x1, 147, 45)
}
func mul(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mul: r=%x\n", r)}()
	return nd(x0, x1, 151, 42)
}
func diw(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("diw: r=%x\n", r)}()
	return nd(x0, x1, 155, 37)
}
func mod(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("mod: r=%x\n", r)}()
	return nd(x0, x1, 23, 7)
}
func adc(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("adc: r=%x\n", r)}()
	MC[x2] = C(C((I(MC[x0]) + I(MC[x1]))))
}
func adi(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("adi: r=%x\n", r)}()
	MI[x2>>2] = I((MI[x0>>2] + MI[x1>>2]))
}
func adf(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("adf: r=%x\n", r)}()
	MF[x2>>3] = F((MF[x0>>3] + MF[x1>>3]))
}
func adz(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("adz: r=%x\n", r)}()
	adf(x0, x1, x2)
	adf((x0 + 8), (x1 + 8), (x2 + 8))
}
func suc(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("suc: r=%x\n", r)}()
	MC[x2] = C(C((I(MC[x0]) - I(MC[x1]))))
}
func sui(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("sui: r=%x\n", r)}()
	MI[x2>>2] = I((MI[x0>>2] - MI[x1>>2]))
}
func suf(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("suf: r=%x\n", r)}()
	MF[x2>>3] = F((MF[x0>>3] - MF[x1>>3]))
}
func suz(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("suz: r=%x\n", r)}()
	suf(x0, x1, x2)
	suf((x0 + 8), (x1 + 8), (x2 + 8))
}
func muc(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("muc: r=%x\n", r)}()
	MC[x2] = C(C((I(MC[x0]) * I(MC[x1]))))
}
func mui(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("mui: r=%x\n", r)}()
	MI[x2>>2] = I((MI[x0>>2] * MI[x1>>2]))
}
func muf(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("muf: r=%x\n", r)}()
	MF[x2>>3] = F((MF[x0>>3] * MF[x1>>3]))
}
func muz(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("muz: r=%x\n", r)}()
	MF[x2>>3] = F(((MF[x0>>3] * MF[x1>>3]) - (MF[(x1+8)>>3] * MF[(x0+8)>>3])))
	MF[(x2+8)>>3] = F(((MF[x0>>3] * MF[(x1+8)>>3]) + (MF[(x0+8)>>3] * MF[x1>>3])))
}
func dic(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("dic: r=%x\n", r)}()
	MC[x2] = C(C((I(MC[x0]) / I(MC[x1]))))
}
func dii(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("dii: r=%x\n", r)}()
	MI[x2>>2] = I((SI(MI[x0>>2]) / SI(MI[x1>>2])))
}
func dif(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("dif: r=%x\n", r)}()
	MF[x2>>3] = F((MF[x0>>3] / MF[x1>>3]))
}
func moi(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("moi: r=%x\n", r)}()
	x0 = I(MI[x0>>2])
	x1 = I(MI[x1>>2])
	MI[x2>>2] = I((SI((x1 + I((SI(x0) % SI(x1))))) % SI(x1)))
}
func diz(x0 I, x1 I, x2 I) {
	//defer func(){fmt.Printf("diz: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8 F
	_, _, _, _, _, _ = x3, x4, x5, x6, x7, x8
	x3 = F(MF[x0>>3])
	x4 = F(MF[(x0+8)>>3])
	x5 = F(MF[x1>>3])
	x6 = F(MF[(x1+8)>>3])
	if math.Abs(x5) >= math.Abs(x6) {
		x7 = F(x6 / x5)
		x8 = F(x5 + (x7 * x6))
		MF[x2>>3] = F(((x3 + (x4 * x7)) / x8))
		MF[(x2+8)>>3] = F(((x4 - (x3 * x7)) / x8))
	} else {
		x7 = F(x5 / x6)
		x8 = F(x6 + (x7 * x5))
		MF[x2>>3] = F(((x4 + (x3 * x7)) / x8))
		MF[(x2+8)>>3] = F((((x4 * x7) - x3) / x8))
	}
}
func abx(x0 I) (r I) {
	//defer func(){fmt.Printf("abx: r=%x\n", r)}()
	return nm(x0, 15, 171)
}
func neg(x0 I) (r I) {
	//defer func(){fmt.Printf("neg: r=%x\n", r)}()
	return nm(x0, 19, 173)
}
func sqr(x0 I) (r I) {
	//defer func(){fmt.Printf("sqr: r=%x\n", r)}()
	return nm(x0, 27, 165)
}
func abc(x0 I, x1 I) {
	//defer func(){fmt.Printf("abc: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(I(MC[x0]))
	if 0 != craz(x2) {
		MC[x1] = C(C((x2 - 32)))
	} else {
		MC[x1] = C(C(x2))
	}
}
func abi(x0 I, x1 I) {
	//defer func(){fmt.Printf("abi: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(MI[x0>>2])
	if SI(x2) < SI(0) {
		MI[x1>>2] = I((0 - x2))
	} else {
		MI[x1>>2] = I(x2)
	}
}
func abf(x0 I, x1 I) {
	//defer func(){fmt.Printf("abf: r=%x\n", r)}()
	MF[x1>>3] = F(math.Abs(MF[x0>>3]))
}
func abz(x0 I, x1 I) {
	//defer func(){fmt.Printf("abz: r=%x\n", r)}()
	MF[x1>>3] = F(hypot(MF[x0>>3], MF[(x0+8)>>3]))
}
func nec(x0 I, x1 I) {
	//defer func(){fmt.Printf("nec: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(I(MC[x0]))
	if 0 != crAZ(x2) {
		MC[x1] = C(C((x2 + 32)))
	} else {
		MC[x1] = C(C(x2))
	}
}
func nei(x0 I, x1 I) {
	//defer func(){fmt.Printf("nei: r=%x\n", r)}()
	MI[x1>>2] = I((0 - MI[x0>>2]))
}
func nef(x0 I, x1 I) {
	//defer func(){fmt.Printf("nef: r=%x\n", r)}()
	MF[x1>>3] = F(-(MF[x0>>3]))
}
func nez(x0 I, x1 I) {
	//defer func(){fmt.Printf("nez: r=%x\n", r)}()
	MF[x1>>3] = F(-(MF[x0>>3]))
	MF[(x1+8)>>3] = F(-(MF[(x0+8)>>3]))
}
func sqc(x0 I, x1 I) {
	//defer func(){fmt.Printf("sqc: r=%x\n", r)}()
	panic("trap")
}
func sqi(x0 I, x1 I) {
	//defer func(){fmt.Printf("sqi: r=%x\n", r)}()
	panic("trap")
}
func sqf(x0 I, x1 I) {
	//defer func(){fmt.Printf("sqf: r=%x\n", r)}()
	MF[x1>>3] = F(math.Sqrt(MF[x0>>3]))
}
func sqz(x0 I, x1 I) {
	//defer func(){fmt.Printf("sqz: r=%x\n", r)}()
	MF[x1>>3] = F(MF[x0>>3])
	MF[(x1+8)>>3] = F(-(MF[(x0+8)>>3]))
}
func lgf(x0 I) (r I) {
	//defer func(){fmt.Printf("lgf: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 != 3 {
		panic("trap")
	}
	x0 = I(use(x0))
	x3 = I(x0 + 8)
	for x4 = 0; x4 < x2; x4++ {
		MF[x3>>3] = F(log(MF[x3>>3]))
		x3 = I(x3 + 8)
	}
	return x0
}
func zre(x0 I) (r I) {
	//defer func(){fmt.Printf("zre: r=%x\n", r)}()
	return zri(x0, 0)
}
func zim(x0 I) (r I) {
	//defer func(){fmt.Printf("zim: r=%x\n", r)}()
	return zri(x0, 8)
}
func zri(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("zri: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(mk(3, x3))
	x6 = I(x5 + 8)
	x4 = I(x4 + x1)
	for x7 = 0; x7 < x3; x7++ {
		MF[x6>>3] = F(MF[x4>>3])
		x6 = I(x6 + 8)
		x4 = I(x4 + 16)
	}
	dx(x0)
	return x5
}
func zan(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("zan: r=%x\n", r)}()
	var x3, x4, x5 I
	_, _, _ = x3, x4, x5
	x3 = I(mk(3, x1))
	x4 = I(x3 + 8)
	for x5 = 0; x5 < x1; x5++ {
		MF[x4>>3] = F(ang(MF[x2>>3], MF[(x2+8)>>3]))
		x2 = I(x2 + 16)
		x4 = I(x4 + 8)
	}
	dx(x0)
	return x3
}
func crAZ(x0 I) (r I) {
	//defer func(){fmt.Printf("crAZ: r=%x\n", r)}()
	if x0 > 64 {
		if x0 < 91 {
			return 1
		}
	}
	return 0
}
func craz(x0 I) (r I) {
	//defer func(){fmt.Printf("craz: r=%x\n", r)}()
	if x0 > 96 {
		if x0 < 123 {
			return 1
		}
	}
	return 0
}
func drv(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("drv: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(mk(0, 2))
	MI[(x2+8)>>2] = I(x0)
	MI[(x2+12)>>2] = I(x1)
	return x2
}
func ecv(x0 I) (r I) {
	//defer func(){fmt.Printf("ecv: r=%x\n", r)}()
	return drv(40, x0)
}
func epv(x0 I) (r I) {
	//defer func(){fmt.Printf("epv: r=%x\n", r)}()
	return drv(41, x0)
}
func ovv(x0 I) (r I) {
	//defer func(){fmt.Printf("ovv: r=%x\n", r)}()
	return drv(123, x0)
}
func riv(x0 I) (r I) {
	//defer func(){fmt.Printf("riv: r=%x\n", r)}()
	return drv(125, x0)
}
func scv(x0 I) (r I) {
	//defer func(){fmt.Printf("scv: r=%x\n", r)}()
	return drv(91, x0)
}
func liv(x0 I) (r I) {
	//defer func(){fmt.Printf("liv: r=%x\n", r)}()
	return drv(93, x0)
}
func ech(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ech: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9
	if 0 != tp(x1) {
		return bin(x1, x0)
	}
	if 7 == tp(x0) {
		rld(x0)
		x2 = I(MI[(x0+8)>>2])
		x3 = I(MI[(x0+12)>>2])
		return mkd(x2, ech(x3, x1))
	}
	x0 = I(lx(x0))
	x4 = I(tp(x0))
	x5 = I(nn(x0))
	x6 = I(8 + x0)
	x7 = I(mk(6, x5))
	x8 = I(x7 + 8)
	rl(x0)
	if x1 < 120 {
		x1 = I(x1 + 128)
	}
	for x9 = 0; x9 < x5; x9++ {
		rx(x1)
		MI[x8>>2] = I(atx(x1, MI[x6>>2]))
		x6 = I(x6 + 4)
		x8 = I(x8 + 4)
	}
	dx(x0)
	dx(x1)
	return x7
}
func ecp(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ecp: r=%x\n", r)}()
	var x2 I
	_ = x2
	rx(x0)
	x2 = I(fst(x0))
	return epi(x2, x0, x1)
}
func epi(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("epi: r=%x\n", r)}()
	var x3, x4, x5, x6, x7 I
	_, _, _, _, _ = x3, x4, x5, x6, x7
	x3 = I(nn(x1))
	if 0 == x3 {
		dx(x0)
		dx(x2)
		return x1
	}
	rxn(x1, x3)
	rxn(x2, x3)
	x4 = I(mk(6, x3))
	x5 = I(x4 + 8)
	for x7 = 0; x7 < x3; x7++ {
		x6 = I(atx(x1, mki(x7)))
		rx(x6)
		MI[x5>>2] = I(cal(x2, l2(x6, x0)))
		x0 = I(x6)
		x5 = I(x5 + 4)
	}
	dx(x6)
	dx(x1)
	dx(x2)
	return x4
}
func ovr(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ovr: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(tp(x1))
	if 2 == x2 {
		return mod(x0, x1)
	}
	return ovs(x0, x1, 0, 0)
}
func scn(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("scn: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(tp(x1))
	if 0 != x2 {
		if x2 < 5 {
			return diw(x0, x1)
		}
	}
	return ovs(x0, x1, enl(mk(6, 0)), 0)
}
func ovi(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("ovi: r=%x\n", r)}()
	return ovs(x1, x2, 0, x0)
}
func sci(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("sci: r=%x\n", r)}()
	return ovs(x1, x2, enl(mk(6, 0)), x0)
}
func scl(x0 I, x1 I) {
	//defer func(){fmt.Printf("scl: r=%x\n", r)}()
	var x2 I
	_ = x2
	if 0 != x0 {
		rx(x1)
		x2 = I(x0 + 8)
		MI[x2>>2] = I(lcat(MI[x2>>2], x1))
	}
}
func ovs(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("ovs: r=%x\n", r)}()
	var x4, x5, x6, x7 I
	_, _, _, _ = x4, x5, x6, x7
	x4 = I(nn(x0))
	rxn(x0, x4)
	x5 = I(x3)
	x6 = I(1)
	if 0 == x5 {
		x5 = I(fst(x0))
		x6 = I(0)
		x4 = I(x4 - 1)
		scl(x2, x5)
	}
	rxn(x1, x4)
	for x7 = 0; x7 < x4; x7++ {
		x5 = I(cal(x1, l2(x5, atx(x0, mki((x7+(1-x6)))))))
		scl(x2, x5)
	}
	dx(x0)
	dx(x1)
	if 0 == x2 {
		return x5
	}
	dx(x5)
	return fst(x2)
}
func fxp(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("fxp: r=%x\n", r)}()
	var x3, x4 I
	_, _ = x3, x4
	x3 = I(x0)
	rx(x0)
	for {
		rx(x0)
		rx(x1)
		x4 = I(atx(x1, x0))
		if 0 != (match(x4, x0) + match(x4, x3)) {
			dx(x0)
			dx(x1)
			dx(x3)
			if 0 != x2 {
				x4 = I(lcat(fst(x2), x4))
			}
			return x4
		}
		scl(x2, x0)
		dx(x0)
		x0 = I(x4)
	}
	return x0
}
func ecr(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("ecr: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8 I
	_, _, _, _, _, _ = x3, x4, x5, x6, x7, x8
	if 7 == tp(x1) {
		rld(x1)
		x3 = I(MI[(x1+8)>>2])
		x4 = I(MI[(x1+12)>>2])
		return mkd(x3, ecr(x0, x4, x2))
	}
	x5 = I(nn(x1))
	x6 = I(mk(6, x5))
	x7 = I(x6 + 8)
	rxn(x0, x5)
	rxn(x1, x5)
	rxn(x2, x5)
	for x8 = 0; x8 < x5; x8++ {
		MI[x7>>2] = I(cal(x2, l2(x0, atx(x1, mki(x8)))))
		x7 = I(x7 + 4)
	}
	dx(x2)
	dx(x0)
	dx(x1)
	return x6
}
func ecl(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("ecl: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8 I
	_, _, _, _, _, _ = x3, x4, x5, x6, x7, x8
	if 7 == tp(x0) {
		rld(x0)
		x3 = I(MI[(x0+8)>>2])
		x4 = I(MI[(x0+12)>>2])
		return mkd(x3, ecl(x4, x1, x2))
	}
	x5 = I(nn(x0))
	x6 = I(mk(6, x5))
	x7 = I(x6 + 8)
	rxn(x0, x5)
	rxn(x1, x5)
	rxn(x2, x5)
	for x8 = 0; x8 < x5; x8++ {
		MI[x7>>2] = I(cal(x2, l2(atx(x0, mki(x8)), x1)))
		x7 = I(x7 + 4)
	}
	dx(x2)
	dx(x0)
	dx(x1)
	return x6
}
func whl(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("whl: r=%x\n", r)}()
	var x4, x5, x6 I
	_, _, _ = x4, x5, x6
	x4 = I(tp(x0))
	if 0 != x4 {
		if 0 != (n32(i32b((x4 == 2))) + n32(i32b((1 == nn(x0))))) {
			panic("trap")
		}
		dx(x0)
		return nlp(x1, x2, x3, MI[(x0+8)>>2])
	}
	x5 = I(x1)
	scl(x3, x5)
	x6 = I(mki(0))
	for {
		rx(x0)
		rx(x2)
		x5 = I(atx(x2, x5))
		scl(x3, x5)
		rx(x5)
		x4 = I(atx(x0, x5))
		if 0 != match(x4, x6) {
			dx(x4)
			dx(x6)
			dx(x2)
			dx(x0)
			if 0 != x3 {
				dx(x5)
				x5 = I(fst(x3))
			}
			return x5
		}
		dx(x4)
	}
	return x0
}
func bin(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("bin: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10, x11 I
	_, _, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10, x11
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x2 != x5 {
		panic("trap")
	}
	x8 = I(mk(2, x6))
	x9 = I(x8 + 8)
	x10 = I(I(MC[x2]))
	for x11 = 0; x11 < x6; x11++ {
		MI[x9>>2] = I(ibin(x4, x7, x3, x2))
		x9 = I(x9 + 4)
		x7 = I(x7 + x10)
	}
	dx(x0)
	dx(x1)
	return x8
}
func ibin(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("ibin: r=%x\n", r)}()
	var x4, x5, x6, x7 I
	_, _, _, _ = x4, x5, x6, x7
	x4 = I(0)
	x5 = I(x2 - 1)
	x6 = I(I(MC[x3]))
	for {
		if SI(x4) > SI(x5) {
			return (x4 - 1)
		}
		x7 = I((x4 + x5) >> 1)
		if 0 != MT[x3].(func(I, I) I)((x0+(x6*x7)), x1) {
			x5 = I(x7 - 1)
		} else {
			x4 = I(x7 + 1)
		}
	}
	return x0
}
func nlp(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("nlp: r=%x\n", r)}()
	var x4, x5 I
	_, _ = x4, x5
	if x3 < 0 {
		panic("trap")
	}
	x4 = I(x0)
	rxn(x1, x3)
	scl(x2, x0)
	for x5 = 0; x5 < x3; x5++ {
		x4 = I(atx(x1, x4))
		scl(x2, x4)
	}
	dx(x1)
	if 0 != x2 {
		dx(x4)
		x4 = I(fst(x2))
	}
	return x4
}
func ecd(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("ecd: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13 I
	_, _, _, _, _, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	x6 = I(tp(x1))
	x7 = I(nn(x1))
	x8 = I(8 + x1)
	if x4 != x7 {
		if x4 == 1 {
			x0 = I(take(x0, x7))
			x4 = I(x7)
			x5 = I(x0 + 8)
		}
		if x7 == 1 {
			x1 = I(take(x1, x4))
			x7 = I(x4)
			x8 = I(x1 + 8)
		}
	}
	x9 = I(nn(x0))
	x10 = I(mk(6, x9))
	x11 = I(x10 + 8)
	rxn(x0, x9)
	rxn(x1, x9)
	rxn(x2, x9)
	for x13 = 0; x13 < x9; x13++ {
		x12 = I(mki(x13))
		rx(x12)
		MI[x11>>2] = I(cal(x2, l2(atx(x0, x12), atx(x1, x12))))
		x11 = I(x11 + 4)
	}
	dx(x2)
	dx(x0)
	dx(x1)
	return x10
}
func val(x0 I) (r I) {
	//defer func(){fmt.Printf("val: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	switch x1 {
	case 0:
		if x0 < 256 {
			return x0
		}
		rl(x0)
		x4 = I(mk(6, x2))
		mv((x4 + 8), (x0 + 8), (4 * x2))
		if x2 == 4 {
			MI[(x4+20)>>2] = I(mki(MI[(x4+20)>>2]))
		}
		dx(x0)
	case 1:
		x4 = I(prs(x0))
		x5 = I(i32b((58 == MI[(x4+8)>>2])))
		x4 = I(evl(x4))
		if 0 != x5 {
			dx(x4)
			x4 = I(0)
		}
	case 5:
		x4 = I(lup(x0))
	case 6:
		x4 = I(evl(x0))
	case 7:
		x4 = I(MI[(x0+12)>>2])
		rx(x4)
		dx(x0)
	default:
		panic("trap")
	}
	return x4
}
func lup(x0 I) (r I) {
	//defer func(){fmt.Printf("lup: r=%x\n", r)}()
	var x1 I
	_ = x1
	x1 = I(MI[(MI[136>>2]+(8+(4*MI[(x0+8)>>2])))>>2])
	rx(x1)
	dx(x0)
	return x1
}
func asn(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("asn: r=%x\n", r)}()
	var x2 I
	_ = x2
	if 5 != tp(x0) {
		panic("trap")
	}
	x2 = I(MI[136>>2] + (8 + (4 * MI[(x0+8)>>2])))
	dx(MI[x2>>2])
	MI[x2>>2] = I(x1)
	rx(x1)
	dx(x0)
	return x1
}
func asd(x0 I) (r I) {
	//defer func(){fmt.Printf("asd: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	rld(x0)
	x1 = I(MI[(x0+8)>>2])
	x2 = I(MI[(x0+12)>>2])
	x3 = I(MI[(x0+16)>>2])
	x4 = I(MI[(x0+20)>>2])
	if x1 != 58 {
		rx(x2)
		x5 = I(lup(x2))
		if 0 != x3 {
			rx(x3)
			x5 = I(atx(x5, x3))
		}
		x4 = I(cal(x1, l2(x5, x4)))
	}
	x5 = I(x4)
	rx(x5)
	if 0 != x3 {
		rx(x2)
		x4 = I(asi(lup(x2), x3, x4))
	}
	dx(asn(x2, x4))
	return x5
}
func asi(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("asi: r=%x\n", r)}()
	var x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16, x17, x18, x19, x20, x21, x22, x23 I
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16, x17, x18, x19, x20, x21, x22, x23
	x3 = I(tp(x0))
	x4 = I(nn(x0))
	x5 = I(8 + x0)
	x6 = I(tp(x1))
	x7 = I(nn(x1))
	x8 = I(8 + x1)
	if x3 == 7 {
		if x6 < 6 {
			rld(x0)
			x9 = I(MI[(x0+8)>>2])
			x10 = I(MI[(x0+12)>>2])
			if x6 == 5 {
				rx(x9)
				x1 = I(fnd(x9, x1))
			}
			return mkd(x9, asi(x10, x1, x2))
		}
	}
	if x6 == 6 {
		if x3 == 7 {
			rld(x0)
			x9 = I(MI[(x0+8)>>2])
			x10 = I(MI[(x0+12)>>2])
			rx(x1)
			x11 = I(fst(x1))
			if 0 != x11 {
				if 5 == tp(x11) {
					rx(x9)
					x11 = I(fnd(x9, x11))
				}
			} else {
				x11 = I(seq(0, nn(x9), 1))
			}
			return mkd(x9, asi(x10, cat(enl(x11), drop(x1, 1)), x2))
		}
		if 0 != (n32(i32b((x3 == 6))) + n32(i32b((x6 == 6)))) {
			panic("trap")
		}
		x12 = I(take(x0, x4))
		x13 = I(x12 + 8)
		rx(x1)
		x14 = I(fst(x1))
		x1 = I(drop(x1, 1))
		if 1 == nn(x1) {
			x1 = I(fst(x1))
		}
		if 0 == x14 {
			x14 = I(seq(0, x4, 1))
		}
		if 2 != tp(x14) {
			panic("trap")
		}
		x15 = I(nn(x14))
		x16 = I(x14 + 8)
		if x15 == 1 {
			dx(x14)
			x17 = I(x13 + (4 * MI[x16>>2]))
			MI[x17>>2] = I(asi(MI[x17>>2], x1, x2))
			return x12
		}
		if x7 != 2 {
			panic("trap")
		}
		if 6 != tp(x2) {
			x2 = I(take(enl(x2), x15))
		}
		if x15 != nn(x2) {
			panic("trap")
		}
		rxn(x1, (x15 - 1))
		rl(x2)
		x18 = I(x2 + 8)
		for x23 = 0; x23 < x15; x23++ {
			x17 = I(x13 + (4 * MI[x16>>2]))
			MI[x17>>2] = I(asi(MI[x17>>2], x1, MI[x18>>2]))
			x16 = I(x16 + 4)
			x18 = I(x18 + 4)
		}
		dx(x14)
		dx(x2)
		return x12
	}
	if x6 != 2 {
		panic("trap")
	}
	x19 = I(tp(x2))
	x20 = I(nn(x2))
	x18 = I(8 + x2)
	if x7 > 1 {
		if x20 == 1 {
			if x20 != x7 {
				if x20 != 1 {
					panic("trap")
				}
				x2 = I(take(x2, x7))
				x20 = I(x7)
				x18 = I(x2 + 8)
			}
		}
	}
	if x3 < 6 {
		if x19 != x3 {
			panic("trap")
		}
		x12 = I(use(x0))
		x13 = I(x12 + 8)
		x21 = I(I(MC[x3]))
		for x23 = 0; x23 < x7; x23++ {
			x9 = I(MI[x8>>2])
			mv((x13 + (x21 * x9)), x18, x21)
			x8 = I(x8 + 4)
			x18 = I(x18 + x21)
		}
		dx(x1)
		dx(x2)
		return x12
	}
	if 0 != (i32b((x3 == 6)) + (i32b((x3 == 5)) * i32b((x19 == 5)))) {
		x12 = I(take(x0, x4))
		if 6 == x3 {
			if 1 == x4 {
				x12 = I(enl(x12))
			}
			if 1 == x7 {
				x2 = I(enl(x2))
				x20 = I(1)
				x19 = I(6)
			}
			if 6 != x19 {
				x2 = I(lx(x2))
			}
		}
		x13 = I(x12 + 8)
		if x7 != x20 {
			panic("trap")
		}
		x18 = I(x2 + 8)
		rl(x2)
		for x23 = 0; x23 < x7; x23++ {
			x9 = I(MI[x8>>2])
			if 0 == i32b((x9 < x4)) {
				panic("trap")
			}
			x22 = I(x13 + (4 * x9))
			dx(MI[x22>>2])
			MI[x22>>2] = I(MI[x18>>2])
			x8 = I(x8 + 4)
			x18 = I(x18 + 4)
		}
		dx(x1)
		dx(x2)
		return x12
	}
	panic("trap")
	return x0
}
func swc(x0 I) (r I) {
	//defer func(){fmt.Printf("swc: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	x4 = I(1)
	for x4 < x2 {
		x5 = I(MI[(x3+(4*x4))>>2])
		rx(x5)
		x5 = I(evl(x5))
		if 0 != (n32((x4 % 2)) | i32b((x4 == (x2 - 1)))) {
			dx(x0)
			return x5
		}
		dx(x5)
		x4 = I(x4 + 1)
		if 0 == MI[(x5+8)>>2] {
			x4 = I(x4 + 1)
		}
	}
	dx(x0)
	return 0
}
func ras(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ras: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _ = x2, x3, x4, x5, x6, x7
	x2 = I(MI[(x0+8)>>2])
	if x1 == 3 {
		if x2 < 256 {
			if 0 != (i32b((x2 == 58)) + i32b((x2 > 128))) {
				if x2 > 128 {
					x2 = I(x2 - 128)
				}
				x3 = I(MI[(x0+12)>>2])
				rxn(x3, 2)
				x4 = I(fst(x3))
				x5 = I(drop(x3, 1))
				if 0 != nn(x5) {
					x5 = I(ltr(x5))
					x6 = I(nn(x5))
					if x6 == 1 {
						x5 = I(fst(x5))
					}
				} else {
					dx(x5)
					x5 = I(0)
				}
				x7 = I(MI[(x0+16)>>2])
				rx(x7)
				dx(x0)
				return lcat(l3(x2, x4, x5), evl(x7))
			}
		}
	}
	return 0
}
func ltr(x0 I) (r I) {
	//defer func(){fmt.Printf("ltr: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 != 6 {
		return x0
	}
	rl(x0)
	x4 = I(mk(6, x2))
	x5 = I(x4 + 8)
	for x6 = 0; x6 < x2; x6++ {
		MI[x5>>2] = I(evl(MI[x3>>2]))
		x5 = I(x5 + 4)
		x3 = I(x3 + 4)
	}
	dx(x0)
	return x4
}
func rtl(x0 I) (r I) {
	//defer func(){fmt.Printf("rtl: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 != 6 {
		return x0
	}
	rl(x0)
	x4 = I(mk(6, x2))
	x5 = I(x4 + (8 + (4 * x2)))
	x3 = I(x3 + (4 * x2))
	for x6 = 0; x6 < x2; x6++ {
		x5 = I(x5 - 4)
		x3 = I(x3 - 4)
		MI[x5>>2] = I(evl(MI[x3>>2]))
	}
	dx(x0)
	return x4
}
func evl(x0 I) (r I) {
	//defer func(){fmt.Printf("evl: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 != 6 {
		if x1 == 5 {
			if x2 == 1 {
				return lup(x0)
			}
		}
		return x0
	}
	if 0 == x2 {
		return x0
	}
	if x2 == 1 {
		return rtl(fst(x0))
	}
	x4 = I(MI[x3>>2])
	if x4 == 36 {
		if x2 > 3 {
			return swc(x0)
		}
	}
	x5 = I(ras(x0, x2))
	if 0 != x5 {
		return asd(x5)
	}
	if x4 == 128 {
		return lst(ltr(x0))
	}
	x0 = I(rtl(x0))
	x2 = I(nn(x0))
	x3 = I(x0 + 8)
	if x4 == 64 {
		if x2 == 4 {
			rl(x0)
			x5 = I(asi(MI[(x0+12)>>2], MI[(x0+16)>>2], MI[(x0+20)>>2]))
			dx(x0)
			return x5
		}
	}
	if x2 == 2 {
		rl(x0)
		x5 = I(atx(MI[x3>>2], MI[(x3+4)>>2]))
		dx(x0)
		return x5
	}
	x6 = I(fnl((x3 + 4), (x2 - 1)))
	if 0 != x6 {
		rx(MI[(x0+8)>>2])
		return prj(MI[(x0+8)>>2], drop(x0, 1), x6)
	}
	rx(MI[x3>>2])
	return cal(MI[x3>>2], drop(x0, 1))
}
func prj(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("prj: r=%x\n", r)}()
	var x3 I
	_ = x3
	x3 = I(mk(0, 3))
	MI[(x3+8)>>2] = I(x0)
	MI[(x3+12)>>2] = I(x1)
	MI[(x3+16)>>2] = I(x2)
	return x3
}
func fnl(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("fnl: r=%x\n", r)}()
	var x2, x3 I
	_, _ = x2, x3
	x2 = I(0)
	for x3 = 0; x3 < x1; x3++ {
		if 0 == MI[x0>>2] {
			if 0 == x2 {
				x2 = I(mk(2, 0))
			}
			x2 = I(ucat(x2, mki(x3)))
		}
		x0 = I(x0 + 4)
	}
	return x2
}
func uqg(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("uqg: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9, x10 I
	_, _, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9, x10
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(mk(x2, 0))
	x6 = I(0)
	x7 = I(I(MC[x2]))
	for x10 = 0; x10 < x3; x10++ {
		x8 = I(fnx(x5, x4))
		if x8 == x6 {
			rx(x0)
			x5 = I(cat(x5, atx(x0, mki(x10))))
			if 0 != x1 {
				x1 = I(lcat(x1, mk(2, 0)))
			}
			x6 = I(x6 + 1)
		}
		if 0 != x1 {
			x9 = I(x1 + (8 + (4 * x8)))
			MI[x9>>2] = I(cat(MI[x9>>2], mki(x10)))
		}
		x4 = I(x4 + x7)
	}
	if 0 != x1 {
		x5 = I(l2(x5, x1))
	}
	dx(x0)
	return x5
}
func unq(x0 I) (r I) {
	//defer func(){fmt.Printf("unq: r=%x\n", r)}()
	return uqg(x0, 0)
}
func grp(x0 I) (r I) {
	//defer func(){fmt.Printf("grp: r=%x\n", r)}()
	return uqg(x0, mk(6, 0))
}
func flr(x0 I) (r I) {
	//defer func(){fmt.Printf("flr: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 > 5 {
		return ech(x0, 223)
	}
	if 0 == x1 {
		dx(x0)
		return mki(x0)
	}
	if 1 == x1 {
		dx(x0)
		return I(MC[x3])
	}
	if 2 == x1 {
		x4 = I(mk(1, x2))
		x5 = I(x4 + 8)
		for x6 = 0; x6 < x2; x6++ {
			MC[(x5 + x6)] = C(C(MI[x3>>2]))
			x3 = I(x3 + 4)
		}
		dx(x0)
		return x4
	}
	if x1 == 3 {
		x4 = I(mk(2, x2))
		x5 = I(x4 + 8)
		for x6 = 0; x6 < x2; x6++ {
			MI[x5>>2] = I(I(MF[x3>>3]))
			x3 = I(x3 + 8)
			x5 = I(x5 + 4)
		}
		dx(x0)
		return x4
	}
	if x1 == 4 {
		return zre(x0)
	}
	panic("trap")
	return x0
}
func ang(x0 F, x1 F) (r F) {
	//defer func(){fmt.Printf("ang: r=%x\n", r)}()
	var x2 F
	_ = x2
	x2 = F(57.29577951308232 * atan2(x1, x0))
	if x2 < 0.0 {
		x2 = F(x2 + 360.0)
	}
	return x2
}
func cst(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cst: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8, x9 I
	_, _, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8, x9
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	if x2 == 5 {
		if x5 == 1 {
			dx(x0)
			return sc(x1)
		}
	}
	if 0 != (n32(i32b((x2 == 2))) + n32(i32b((x3 == 1)))) {
		panic("trap")
	}
	dx(x0)
	x0 = I(MI[(x0+8)>>2])
	if SI(x0) < SI(0) {
		x0 = I(-(x0))
		x8 = I(x6 / I(MC[x0]))
		if x6 != (x8 * I(MC[x0])) {
			panic("trap")
		}
		x9 = I(use(x1))
		MI[x9>>2] = I((x8 | (x0 << 29)))
		return x9
	}
	if 0 == x6 {
		dx(x1)
		if 7 == x0 {
			return mkd(mk(5, 0), mk(6, 0))
		}
		return mk(x0, 0)
	}
	if 0 != (i32b((x5 > x0)) + i32b((x5 > 4))) {
		panic("trap")
	}
	if 8 == x0 {
		x8 = I(x6 * I(MC[x5]))
		x9 = I(use(x1))
		MI[x9>>2] = I((x8 | (1 << 29)))
		return x9
	}
	for SI(x5) < SI(x0) {
		x1 = I(up(x1, x5, x6))
		x5 = I(x5 + 1)
	}
	return x1
}
func flp(x0 I) (r I) {
	//defer func(){fmt.Printf("flp: r=%x\n", r)}()
	var x1, x2 I
	_, _ = x1, x2
	x1 = I(nn(MI[(x0+8)>>2]))
	x2 = I(nn(x0))
	return atx(ovr(x0, 44), ecr(mul(mki(x1), seq(0, x2, 1)), seq(0, x1, 1), 43))
}
func rnd(x0 I) (r I) {
	//defer func(){fmt.Printf("rnd: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	if 1073741825 != MI[x0>>2] {
		panic("trap")
	}
	dx(x0)
	x0 = I(MI[(x0+8)>>2])
	x1 = I(mk(2, x0))
	x2 = I(x1 + 8)
	for x3 = 0; x3 < x0; x3++ {
		MI[x2>>2] = I(rng(0))
		x2 = I(x2 + 4)
	}
	return x1
}
func rng(x0 I) (r I) {
	//defer func(){fmt.Printf("rng: r=%x\n", r)}()
	x0 = I(MI[12>>2])
	x0 = I(x0 ^ (x0 << 13))
	x0 = I(x0 ^ (x0 >> 17))
	x0 = I(x0 ^ (x0 << 5))
	MI[12>>2] = I(x0)
	return x0
}
func xxx(x0 I) (r I) {
	//defer func(){fmt.Printf("xxx: r=%x\n", r)}()
	panic("trap")
	return x0
}
func drw(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("drw: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8 I
	_, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	x5 = I(tp(x1))
	x6 = I(nn(x1))
	x7 = I(8 + x1)
	x8 = I(MI[x4>>2])
	if x5 == 7 {
		draw(x8, MI[(x4+4)>>2], x1)
	}
	if x5 == 2 {
		draw(x8, (x6 / x8), x1)
	}
	dx(x0)
	dx(x1)
	return 0
}
func sadv(x0 I) (r I) {
	//defer func(){fmt.Printf("sadv: r=%x\n", r)}()
	var x1 I
	_ = x1
	if x0 == 39 {
		x1 = I(1)
	} else if x0 == 47 {
		x1 = I(1)
	} else if x0 == 92 {
		x1 = I(1)
	} else {
		x1 = I(0)
	}
	return x1
}
func out(x0 I) (r I) {
	//defer func(){fmt.Printf("out: r=%x\n", r)}()
	var x1 I
	_ = x1
	rx(x0)
	x1 = I(x0)
	if 1 != tp(x1) {
		x1 = I(kst(x0))
	}
	printc((x1 + 8), nn(x1))
	dx(x1)
	return x0
}
func kst(x0 I) (r I) {
	//defer func(){fmt.Printf("kst: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	if 0 == nn(x0) {
		if x1 > 1 {
			if x1 < 6 {
				dx(x0)
				x2 = I(cc(cc(mkc(48), 35), 48))
				if x1 == 3 {
					x2 = I(cc(x2, 46))
				}
				if x1 == 4 {
					x2 = I(cc(x2, 97))
				}
				if x1 == 5 {
					MC[(x2 + 10)] = C(C(96))
				}
				return x2
			}
		}
	}
	if 7 == x1 {
		rld(x0)
		x3 = I(MI[(x0+8)>>2])
		x4 = I(MI[(x0+12)>>2])
		x3 = I(cc(kst(x3), 33))
		if 0 == nn(x4) {
			dx(x4)
			x4 = I(mki(0))
		}
		return ucat(x3, kst(x4))
	}
	if 6 == x1 {
		if 1 == nn(x0) {
			return ucat(mkc(44), kst(fst(x0)))
		}
		x0 = I(ech(x0, 235))
	} else {
		x0 = I(str(x0))
	}
	switch x1 {
	case 0:
		x2 = I(x0)
	case 1:
		x2 = I(cc(ucat(mkc(34), x0), 34))
	case 5:
		x2 = I(ucat(mkc(96), jon(x0, mkc(96))))
	case 6:
		x2 = I(cc(ucat(mkc(40), jon(x0, mkc(59))), 41))
	default:
		x2 = I(jon(x0, mkc(32)))
	}
	return x2
}
func str(x0 I) (r I) {
	//defer func(){fmt.Printf("str: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 == 1 {
		return x0
	}
	if 0 == x1 {
		return cg(x0, x2)
	}
	if 0 != (i32b((x1 > 5)) + n32(i32b((x2 == 1)))) {
		return ech(x0, 164)
	}
	switch x1 {
	case 2:
		x4 = I(ci(MI[x3>>2]))
	case 3:
		x4 = I(cf(MF[x3>>3]))
	case 4:
		x4 = I(cz(MF[x3>>3], MF[(x3+8)>>3]))
	case 5:
		rx(x0)
		x4 = I(cs(x0))
	default:
		panic("trap")
	}
	dx(x0)
	return x4
}
func cc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cc: r=%x\n", r)}()
	var x2 I
	_ = x2
	x2 = I(nn(x0))
	if bk(1, x2) < bk(1, (x2+1)) {
		return ucat(x0, mkc(x1))
	}
	MC[(x0 + (8 + x2))] = C(C(x1))
	MI[x0>>2] = I((1 + MI[x0>>2]))
	return x0
}
func ng(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ng: r=%x\n", r)}()
	if 0 != x1 {
		x0 = I(ucat(mkc(45), x0))
	}
	return x0
}
func cg(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("cg: r=%x\n", r)}()
	var x2 I
	_ = x2
	if 0 != (n32(x0) + i32b((x0 == 128))) {
		return mk(1, 0)
	}
	if x0 < 127 {
		return mkc(x0)
	}
	if x0 < 256 {
		return cc(mkc((x0 - 128)), 58)
	}
	if x1 == 2 {
		rl(x0)
		x2 = I(cat(str(MI[(x0+12)>>2]), str(MI[(x0+8)>>2])))
	}
	if x1 == 3 {
		rl(x0)
		dx(MI[(x0+16)>>2])
		x2 = I(kst(MI[(x0+12)>>2]))
		MC[(x2 + 8)] = C(C(91))
		MC[(x2 + (7 + nn(x2)))] = C(C(93))
		x2 = I(ucat(str(MI[(x0+8)>>2]), x2))
	}
	if x1 == 4 {
		x2 = I(MI[(x0+8)>>2])
		rx(x2)
	}
	dx(x0)
	return x2
}
func ci(x0 I) (r I) {
	//defer func(){fmt.Printf("ci: r=%x\n", r)}()
	var x1, x2, x3 I
	_, _, _ = x1, x2, x3
	if 0 == x0 {
		return mkc(48)
	}
	x1 = I(0)
	if SI(x0) < SI(0) {
		x0 = I(0 - x0)
		x1 = I(1)
	}
	x2 = I(mk(1, 0))
	for 0 != x0 {
		x3 = I(x0 % 10)
		x2 = I(cc(x2, (48 + x3)))
		x0 = I(x0 / 10)
	}
	if 0 == nn(x2) {
		x2 = I(cc(x2, 48))
	}
	return ng(rev(x2), x1)
}
func cf(x0 F) (r I) {
	//defer func(){fmt.Printf("cf: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6, x7 I
	_, _, _, _, _, _, _ = x1, x2, x3, x4, x5, x6, x7
	if x0 != x0 {
		return cc(mkc(48), 110)
	}
	if x0 == 0.0 {
		return cc(cc(mkc(48), 46), 48)
	}
	x1 = I(0)
	if x0 < 0.0 {
		x1 = I(1)
		x0 = F(-(x0))
	}
	if x0 > 1.7976931348623157e+308 {
		return ng(cc(mkc(48), 119), x1)
	}
	x2 = I(0)
	for x0 > 1000.0 {
		x2 = I(x2 + 3)
		x0 = F(x0 / 1000.0)
	}
	x3 = I(7)
	if x0 < 1.0 {
		x3 = I(x3 + 1)
		if x0 < 0.1 {
			x3 = I(x3 + 1)
			if x0 < 0.01 {
				x3 = I(x3 + 1)
				if x0 < 0.001 {
					x3 = I(7)
					for x0 < 1.0 {
						x2 = I(x2 - 3)
						x0 = F(x0 * 1000.0)
					}
				}
			}
		}
	}
	x4 = I(I(x0))
	x5 = I(ci(x4))
	x0 = F(x0 - F(x4))
	x3 = I(x3 - nn(x5))
	if SI(x3) < SI(1) {
		x3 = I(1)
	}
	x5 = I(cc(x5, 46))
	x6 = I(0)
	for x7 = 0; x7 < x3; x7++ {
		x0 = F(x0 * 10.0)
		x4 = I(I(x0))
		x5 = I(cc(x5, (48 + x4)))
		x0 = F(x0 - F(x4))
		x6 = I((1 + x6) * n32((x4 + n32(x7))))
	}
	x5 = I(drop(x5, -(x6)))
	if 0 != x2 {
		x5 = I(ucat(cc(x5, 101), ci(x2)))
	}
	return ng(x5, x1)
}
func cz(x0 F, x1 F) (r I) {
	//defer func(){fmt.Printf("cz: r=%x\n", r)}()
	var x2 F

	var x3 I
	_, _ = x2, x3
	x2 = F(hypot(x0, x1))
	x3 = I(I((0.5 + ang(x0, x1))))
	return ucat(cc(cf(x2), 97), ci(x3))
}
func prs(x0 I) (r I) {
	//defer func(){fmt.Printf("prs: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if x1 != 1 {
		panic("trap")
	}
	x2 = I(x2 + x3)
	MI[8>>2] = I(x3)
	if 0 != x2 {
		if 47 == I(MC[x3]) {
			MI[8>>2] = I(com(x3, x2))
		}
	}
	x4 = I(sq(x2))
	if 1 == nn(x4) {
		x4 = I(fst(x4))
	} else {
		x4 = I(cat(128, x4))
	}
	dx(x0)
	return x4
}
func sq(x0 I) (r I) {
	//defer func(){fmt.Printf("sq: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(mk(6, 0))
	x2 = I(ex(pt(x0), x0))
	if 0 != x2 {
		x1 = I(lcat(x1, x2))
	}
	for {
		x3 = I(ws(x0))
		x4 = I(MI[8>>2])
		if 0 == x3 {
			x3 = I(I(MC[x4]))
			x3 = I(n32((i32b((x3 == 59)) + i32b((x3 == 10)))))
		}
		if 0 != x3 {
			if x4 < x0 {
				MI[8>>2] = I((1 + x4))
			}
			return x1
		}
		MI[8>>2] = I((1 + x4))
		if 0 == nn(x1) {
			x1 = I(lcat(x1, 0))
		}
		x1 = I(lcat(x1, ex(pt(x0), x0)))
	}
	panic("trap")
	return x0
}
func ex(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("ex: r=%x\n", r)}()
	var x2 I
	_ = x2
	if 0 != (n32(x0) + ws(x1)) {
		return x0
	}
	x2 = I(I(MC[MI[8>>2]]))
	if 0 != (is(x2, 32) + i32b((x2 == 10))) {
		return x0
	}
	x2 = I(pt(x1))
	if 0 != isv(x2) {
		if 0 == isv(x0) {
			return l3(x2, x0, ex(pt(x1), x1))
		}
	}
	return l2(x0, ex(x2, x1))
}
func pt(x0 I) (r I) {
	//defer func(){fmt.Printf("pt: r=%x\n", r)}()
	var x1, x2, x3, x4, x5, x6 I
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6
	x1 = I(tok(x0))
	if 0 == x1 {
		x2 = I(MI[8>>2])
		if x2 == x0 {
			return 0
		}
		x3 = I(i32b((123 == I(MC[x2]))))
		if 0 != (x3 + i32b((40 == I(MC[x2])))) {
			MI[8>>2] = I((1 + x2))
			if 0 != x3 {
				x4 = I(0)
				if 91 == I(MC[(1+x2)]) {
					MI[8>>2] = I((2 + x2))
					x4 = I(sq(x0))
					if 0 == nn(x4) {
						x4 = I(lcat(x4, mk(5, 0)))
					}
					x4 = I(ovr(x4, 44))
				}
				x1 = I(sq(x0))
				x1 = I(lam(x2, MI[8>>2], x1, x4))
			} else {
				x1 = I(sq(x0))
				x5 = I(nn(x1))
				if x5 == 1 {
					x1 = I(fst(x1))
				}
				if x5 > 1 {
					x1 = I(enl(x1))
				}
			}
		}
	}
	for {
		x2 = I(MI[8>>2])
		x6 = I(I(MC[x2]))
		if 0 != (i32b((x2 == x0)) + i32b((32 == I(MC[(x2-1)])))) {
			return x1
		}
		if 0 != is(x6, 16) {
			x1 = I(l2(tok(x0), x1))
		} else if x6 == 91 {
			MI[8>>2] = I((1 + x2))
			x2 = I(sq(x0))
			if 0 == nn(x2) {
				x2 = I(lcat(x2, 0))
			}
			x1 = I(cat(enl(x1), x2))
		} else {
			return x1
		}
	}
	panic("trap")
	return x1
}
func isv(x0 I) (r I) {
	//defer func(){fmt.Printf("isv: r=%x\n", r)}()
	var x1, x2, x3, x4 I
	_, _, _, _ = x1, x2, x3, x4
	x1 = I(tp(x0))
	x2 = I(nn(x0))
	x3 = I(8 + x0)
	if 0 == x1 {
		return 1
	}
	if x1 == 6 {
		if x2 == 2 {
			x4 = I(MI[x3>>2])
			if x4 < 256 {
				if 0 != (is(x4, 16) | is((x4-128), 16)) {
					return 1
				}
			}
		}
	}
	return 0
}
func lac(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("lac: r=%x\n", r)}()
	var x2, x3, x4, x5, x6 I
	_, _, _, _, _ = x2, x3, x4, x5, x6
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	if x2 == 6 {
		if 1 == x3 {
			if 5 == tp(MI[(x0+8)>>2]) {
				return x1
			}
		}
		for x6 = 0; x6 < x3; x6++ {
			x1 = I(lac(MI[x4>>2], x1))
			x4 = I(x4 + 4)
		}
	}
	if x2 == 5 {
		if x3 == 1 {
			x5 = I(MI[x4>>2])
			if x5 > x1 {
				if x5 < 4 {
					return x5
				}
			}
		}
	}
	return x1
}
func loc(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("loc: r=%x\n", r)}()
	var x2, x3, x4, x5, x6, x7, x8 I
	_, _, _, _, _, _, _ = x2, x3, x4, x5, x6, x7, x8
	x2 = I(tp(x0))
	x3 = I(nn(x0))
	x4 = I(8 + x0)
	if x2 != 6 {
		return x1
	}
	for x8 = 0; x8 < x3; x8++ {
		x1 = I(loc(MI[x4>>2], x1))
		x4 = I(x4 + 4)
	}
	x4 = I(x0 + 8)
	if x3 == 3 {
		if 58 == MI[x4>>2] {
			x5 = I(MI[(x4+4)>>2])
			rx(x5)
			x6 = I(fst(x5))
			x7 = I(nn(x1))
			if x7 == fnx(x1, (x6+8)) {
				rx(x6)
				x1 = I(cat(x1, x6))
			}
			dx(x6)
		}
	}
	return x1
}
func lam(x0 I, x1 I, x2 I, x3 I) (r I) {
	//defer func(){fmt.Printf("lam: r=%x\n", r)}()
	var x4, x5, x6, x7 I
	_, _, _, _ = x4, x5, x6, x7
	if 1 == nn(x2) {
		x2 = I(fst(x2))
	} else {
		x2 = I(cat(128, x2))
	}
	if 0 == x3 {
		x4 = I(MI[148>>2])
		rx(x4)
		x3 = I(take(x4, lac(x2, 0)))
	}
	x5 = I(nn(x3))
	x3 = I(loc(x2, x3))
	x6 = I(x1 - x0)
	x7 = I(mk(1, x6))
	mv((x7 + 8), x0, x6)
	x4 = I(mk(0, 4))
	MI[(x4+8)>>2] = I(x7)
	MI[(x4+12)>>2] = I(x2)
	MI[(x4+16)>>2] = I(x3)
	MI[(x4+20)>>2] = I(x5)
	return x4
}
func ws(x0 I) (r I) {
	//defer func(){fmt.Printf("ws: r=%x\n", r)}()
	var x1, x2 I
	_, _ = x1, x2
	x1 = I(MI[8>>2])
	if 47 == I(MC[x1]) {
		x2 = I(I(MC[(x1 - 1)]))
		if 0 != (i32b((x2 == 32)) + i32b((x2 == 10))) {
			x1 = I(com(x1, x0))
		}
	}
	for {
		if x1 == x0 {
			MI[8>>2] = I(x1)
			return 1
		}
		x2 = I(I(MC[x1]))
		if 0 != (i32b((x2 == 10)) + is(x2, 64)) {
			MI[8>>2] = I(x1)
			return 0
		}
		x1 = I(x1 + 1)
		if 47 == I(MC[x1]) {
			x1 = I(com(x1, x0))
		}
	}
	return x0
}
func com(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("com: r=%x\n", r)}()
	for x0 < x1 {
		if 10 == I(MC[x0]) {
			return x0
		}
		x0 = I(x0 + 1)
	}
	return x0
}
func tok(x0 I) (r I) {
	//defer func(){fmt.Printf("tok: r=%x\n", r)}()
	var x1, x2, x3, x4, x5 I
	_, _, _, _, _ = x1, x2, x3, x4, x5
	if 0 != ws(x0) {
		return 0
	}
	x1 = I(MI[8>>2])
	x2 = I(I(MC[x1]))
	if 0 != (is(x2, 32) + i32b((x2 == 10))) {
		return 0
	}
	x4 = 5
	for x5 = 0; x5 < x4; x5++ {
		x3 = I(MT[(x5+136)].(func(I, I, I) I)(x2, x1, x0))
		if 0 != x3 {
			return x3
		}
	}
	return 0
}
func puj(x0 I, x1 I, x2 I) (r J) {
	//defer func(){fmt.Printf("puj: r=%x\n", r)}()
	var x3 J
	_ = x3
	if 0 == is(x0, 4) {
		return 0
	}
	for 0 != (is(x0, 4) * i32b((x1 < x2))) {
		x3 = J(x3 * 10)
		x3 = J(x3 + J((x0 - 48)))
		x1 = I(x1 + 1)
		x0 = I(I(MC[x1]))
	}
	if x3 == 0 {
		if 120 == x0 {
			return 0
		}
	}
	MI[8>>2] = I(x1)
	return x3
}
func pin(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("pin: r=%x\n", r)}()
	var x3 J
	_ = x3
	x3 = J(puj(x0, x1, x2))
	if x1 != MI[8>>2] {
		return mki(I(x3))
	}
	if x0 == 45 {
		x1 = I(x1 + 1)
		if x1 < x2 {
			x0 = I(I(MC[x1]))
			MI[8>>2] = I(x1)
			x3 = J(puj(x0, x1, x2))
			if x1 != MI[8>>2] {
				return mki(-(I(x3)))
			}
			MI[8>>2] = I((x1 - 1))
		}
	}
	return 0
}
func pfl(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("pfl: r=%x\n", r)}()
	var x3, x4, x5, x7, x8, x9, x10, x11 I

	var x6 F
	_, _, _, _, _, _, _, _, _ = x3, x4, x5, x6, x7, x8, x9, x10, x11
	x3 = I(0)
	if x0 == 45 {
		x4 = I(I(MC[(x1 - 1)]))
		if 0 != (i32b((x4 == 34)) + (i32b((x4 == 93)) + (i32b((x4 == 41)) + is(x4, 7)))) {
			return 0
		}
		x3 = I(1)
	}
	x5 = I(pin(x0, x1, x2))
	x1 = I(MI[8>>2])
	if 0 != (i32b((x1 == x2)) + n32(x5)) {
		return x5
	}
	if 46 == I(MC[x1]) {
		x5 = I(up(x5, 2, 1))
		x1 = I(x1 + 1)
		MI[8>>2] = I(x1)
		if x1 < x2 {
			x0 = I(I(MC[x1]))
			x6 = F(F(puj(x0, x1, x2)))
			if x1 != MI[8>>2] {
				x10 = (MI[8>>2] - x1)
				for x11 = 0; x11 < x10; x11++ {
					x6 = F(x6 / 10.0)
				}
				x7 = I(x5 + 8)
				if MF[x7>>3] < 0.0 {
					x6 = F(-(x6))
				}
				MF[x7>>3] = F((MF[x7>>3] + x6))
			}
		}
	}
	x1 = I(MI[8>>2])
	if x1 < x2 {
		if 101 == I(MC[x1]) {
			MI[8>>2] = I((x1 + 1))
			x8 = I(pin(I(MC[(1+x1)]), (1 + x1), x2))
			if 0 == x8 {
				MI[8>>2] = I(x1)
				return x5
			}
			x9 = I(MI[(x8+8)>>2])
			dx(x8)
			x6 = F(MF[(x5+8)>>3])
			for SI(x9) < SI(0) {
				x6 = F(x6 / 10.0)
				x9 = I(x9 + 1)
			}
			for x9 > 0 {
				x6 = F(x6 * 10.0)
				x9 = I(x9 - 1)
			}
			MF[(x5+8)>>3] = F(x6)
		}
	}
	if 0 != x3 {
		x6 = F(MF[(x5+8)>>3])
		if x6 > 0.0 {
			MF[(x5+8)>>3] = F(-(x6))
		}
	}
	return x5
}
func num(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("num: r=%x\n", r)}()
	var x3, x5 I

	var x4 F
	_, _, _ = x3, x4, x5
	x3 = I(pfl(x0, x1, x2))
	if 0 == x3 {
		return x3
	}
	x1 = I(MI[8>>2])
	x0 = I(I(MC[x1]))
	if x1 < x2 {
		if 0 != (i32b((119 == x0)) + (i32b((110 == x0)) + (i32b((112 == x0)) + i32b((97 == x0))))) {
			if 2 == tp(x3) {
				x3 = I(up(x3, 2, 1))
			}
			x1 = I(x1 + 1)
			MI[8>>2] = I(x1)
			if 97 != x0 {
				if 112 == x0 {
					x4 = F(3.141592653589793 * MF[(x3+8)>>3])
				}
				if 110 == x0 {
					x4 = F(NAN)
				}
				if 119 == x0 {
					x4 = F(INFINITY)
				}
				MF[(x3+8)>>3] = F(x4)
				return x3
			}
			x3 = I(up(x3, 3, 1))
			x5 = I(pfl(I(MC[x1]), x1, x2))
			if 0 == x5 {
				x5 = I(mki(0))
			}
			if 2 == tp(x5) {
				x5 = I(up(x5, 2, 1))
			}
			x3 = I(atx(x3, x5))
		}
	}
	return x3
}
func nms(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("nms: r=%x\n", r)}()
	var x3, x4 I
	_, _ = x3, x4
	x3 = I(num(x0, x1, x2))
	if 0 == x3 {
		return x3
	}
	for {
		x1 = I(MI[8>>2])
		x0 = I(I(MC[x1]))
		if (x1 + 2) > x2 {
			return x3
		}
		if x0 != 32 {
			return x3
		}
		x1 = I(x1 + 1)
		MI[8>>2] = I(x1)
		x4 = I(num(I(MC[x1]), x1, x2))
		if 0 == x4 {
			MI[8>>2] = I((x1 - 1))
			return x3
		}
		x3 = I(upx(x3, x4))
		x4 = I(upx(x4, x3))
		x3 = I(cat(x3, x4))
	}
	return x3
}
func vrb(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("vrb: r=%x\n", r)}()
	var x3 I
	_ = x3
	if 0 == is(x0, 24) {
		return 0
	}
	if 32 == I(MC[(x1-1)]) {
		if x0 == 92 {
			MI[8>>2] = I((1 + x1))
			return 160
		}
		if x0 == 39 {
			x1 = I(x1 + 1)
		}
	}
	x3 = I(I(MC[x1]))
	if x2 > (1 + x0) {
		if 58 == I(MC[(1+x1)]) {
			x1 = I(x1 + 1)
			x3 = I(x3 + 128)
		}
	}
	MI[8>>2] = I((1 + x1))
	return x3
}
func chr(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("chr: r=%x\n", r)}()
	var x3, x4, x5 I
	_, _, _ = x3, x4, x5
	if 48 == x0 {
		if 120 == I(MC[(1+x1)]) {
			return phx((2 + x1), x2)
		}
	}
	if x0 != 34 {
		return 0
	}
	x3 = I(1 + x1)
	for {
		x1 = I(x1 + 1)
		if x1 == x2 {
			panic("trap")
		}
		if 34 == I(MC[x1]) {
			x4 = I(x1 - x3)
			x5 = I(mk(1, x4))
			mv((x5 + 8), x3, x4)
			MI[8>>2] = I((1 + x1))
			return x5
		}
	}
	return x5
}
func phx(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("phx: r=%x\n", r)}()
	var x2, x3, x4, x5 I
	_, _, _, _ = x2, x3, x4, x5
	x2 = I(mk(1, 0))
	x3 = I(1)
	for {
		x4 = I(I(MC[x0]))
		if 0 != (i32b((x1 <= x0)) + n32(is(x4, 5))) {
			MI[8>>2] = I(x0)
			return x2
		}
		x4 = I(x4 - ((48 * i32b((x4 < 58))) + (87 * i32b((x4 > 96)))))
		x3 = I(n32(x3))
		if 0 != x3 {
			x2 = I(cc(x2, (x4 + (x5 << 4))))
		}
		x5 = I(x4)
		x0 = I(x0 + 1)
	}
	return x0
}
func nam(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("nam: r=%x\n", r)}()
	var x3, x4, x5 I
	_, _, _ = x3, x4, x5
	if 0 == is(x0, 3) {
		return 0
	}
	x3 = I(x1)
	for {
		x1 = I(x1 + 1)
		if 0 != (i32b((x1 == x2)) + n32(is(I(MC[x1]), 7))) {
			x4 = I(x1 - x3)
			x5 = I(mk(1, x4))
			mv((x5 + 8), x3, x4)
			MI[8>>2] = I(x1)
			return sc(x5)
		}
	}
	return x0
}
func sym(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("sym: r=%x\n", r)}()
	var x3 I
	_ = x3
	if 0 != (i32b((x1 == x2)) + n32(i32b((x0 == 96)))) {
		return 0
	}
	x1 = I(x1 + 1)
	x0 = I(I(MC[x1]))
	MI[8>>2] = I(x1)
	if x1 < x2 {
		x3 = I(nam(x0, x1, x2))
		if 0 != x3 {
			return x3
		}
		x3 = I(chr(x0, x1, x2))
		if 0 != x3 {
			return sc(x3)
		}
	}
	x3 = I(mk(5, 1))
	MI[(x3+8)>>2] = I(0)
	return x3
}
func sms(x0 I, x1 I, x2 I) (r I) {
	//defer func(){fmt.Printf("sms: r=%x\n", r)}()
	var x3, x4 I
	_, _ = x3, x4
	x3 = I(sym(x0, x1, x2))
	if 0 == x3 {
		return x3
	}
	for {
		x1 = I(MI[8>>2])
		x4 = I(sym(I(MC[x1]), x1, x2))
		if 0 == x4 {
			return enl(x3)
		}
		x3 = I(cat(x3, x4))
	}
	return x3
}
func is(x0 I, x1 I) (r I) {
	//defer func(){fmt.Printf("is: r=%x\n", r)}()
	return (x1 & cla(x0))
}
func cla(x0 I) (r I) {
	//defer func(){fmt.Printf("cla: r=%x\n", r)}()
	if 128 < (x0 - 32) {
		return 0
	}
	return I(MC[(128 + x0)])
}

var MT [256]interface{}

func mt_init() {
	MT[0] = xxx
	MT[1] = gtc
	MT[2] = gti
	MT[3] = gtf
	MT[4] = xxx
	MT[5] = gti
	MT[6] = gtl
	MT[7] = mod
	MT[8] = xxx
	MT[9] = eqc
	MT[10] = eqi
	MT[11] = eqf
	MT[12] = eqz
	MT[13] = eqi
	MT[14] = eqL
	MT[15] = xxx
	MT[16] = abc
	MT[17] = abi
	MT[18] = abf
	MT[19] = abz
	MT[20] = nec
	MT[21] = nei
	MT[22] = nef
	MT[23] = nez
	MT[24] = xxx
	MT[25] = moi
	MT[26] = xxx
	MT[27] = xxx
	MT[28] = sqc
	MT[29] = sqi
	MT[30] = sqf
	MT[31] = sqz
	MT[32] = xxx
	MT[33] = mkd
	MT[34] = xxx
	MT[35] = rsh
	MT[36] = cst
	MT[37] = diw
	MT[38] = min
	MT[39] = ecv
	MT[40] = ecd
	MT[41] = epi
	MT[42] = mul
	MT[43] = add
	MT[44] = cat
	MT[45] = sub
	MT[46] = cal
	MT[47] = ovv
	MT[48] = xxx
	MT[49] = xxx
	MT[50] = xxx
	MT[51] = xxx
	MT[52] = xxx
	MT[53] = xxx
	MT[54] = xxx
	MT[55] = xxx
	MT[56] = xxx
	MT[57] = xxx
	MT[58] = asn
	MT[59] = xxx
	MT[60] = les
	MT[61] = eql
	MT[62] = mor
	MT[63] = fnd
	MT[64] = atx
	MT[65] = xxx
	MT[66] = xxx
	MT[67] = xxx
	MT[68] = xxx
	MT[69] = xxx
	MT[70] = nmf
	MT[71] = xxx
	MT[72] = xxx
	MT[73] = xxx
	MT[74] = xxx
	MT[75] = xxx
	MT[76] = xxx
	MT[77] = xxx
	MT[78] = xxx
	MT[79] = xxx
	MT[80] = xxx
	MT[81] = xxx
	MT[82] = xxx
	MT[83] = xxx
	MT[84] = xxx
	MT[85] = xxx
	MT[86] = xxx
	MT[87] = xxx
	MT[88] = xxx
	MT[89] = xxx
	MT[90] = xxx
	MT[91] = sci
	MT[92] = scv
	MT[93] = ecl
	MT[94] = exc
	MT[95] = cut
	MT[96] = xxx
	MT[97] = xxx
	MT[98] = xxx
	MT[99] = xxx
	MT[100] = drw
	MT[101] = xxx
	MT[102] = xxx
	MT[103] = xxx
	MT[104] = xxx
	MT[105] = xxx
	MT[106] = xxx
	MT[107] = xxx
	MT[108] = xxx
	MT[109] = xxx
	MT[110] = xxx
	MT[111] = xxx
	MT[112] = xxx
	MT[113] = xxx
	MT[114] = xxx
	MT[115] = xxx
	MT[116] = xxx
	MT[117] = xxx
	MT[118] = xxx
	MT[119] = xxx
	MT[120] = xxx
	MT[121] = xxx
	MT[122] = xxx
	MT[123] = ovi
	MT[124] = max
	MT[125] = ecr
	MT[126] = mtc
	MT[127] = xxx
	MT[128] = xxx
	MT[129] = sin
	MT[130] = cos
	MT[131] = exp
	MT[132] = log
	MT[133] = xxx
	MT[134] = xxx
	MT[135] = xxx
	MT[136] = chr
	MT[137] = nms
	MT[138] = vrb
	MT[139] = nam
	MT[140] = sms
	MT[141] = xxx
	MT[142] = xxx
	MT[143] = xxx
	MT[144] = adc
	MT[145] = adi
	MT[146] = adf
	MT[147] = adz
	MT[148] = suc
	MT[149] = sui
	MT[150] = suf
	MT[151] = suz
	MT[152] = muc
	MT[153] = mui
	MT[154] = muf
	MT[155] = muz
	MT[156] = dic
	MT[157] = dii
	MT[158] = dif
	MT[159] = diz
	MT[160] = out
	MT[161] = til
	MT[162] = xxx
	MT[163] = cnt
	MT[164] = str
	MT[165] = sqr
	MT[166] = wer
	MT[167] = epv
	MT[168] = ech
	MT[169] = ecp
	MT[170] = fst
	MT[171] = abx
	MT[172] = enl
	MT[173] = neg
	MT[174] = val
	MT[175] = riv
	MT[176] = xxx
	MT[177] = xxx
	MT[178] = xxx
	MT[179] = xxx
	MT[180] = xxx
	MT[181] = xxx
	MT[182] = xxx
	MT[183] = xxx
	MT[184] = xxx
	MT[185] = xxx
	MT[186] = lst
	MT[187] = xxx
	MT[188] = grd
	MT[189] = grp
	MT[190] = gdn
	MT[191] = unq
	MT[192] = typ
	MT[193] = xxx
	MT[194] = xxx
	MT[195] = xxx
	MT[196] = xxx
	MT[197] = xxx
	MT[198] = xxx
	MT[199] = xxx
	MT[200] = xxx
	MT[201] = xxx
	MT[202] = xxx
	MT[203] = xxx
	MT[204] = xxx
	MT[205] = xxx
	MT[206] = xxx
	MT[207] = xxx
	MT[208] = xxx
	MT[209] = xxx
	MT[210] = xxx
	MT[211] = xxx
	MT[212] = xxx
	MT[213] = xxx
	MT[214] = xxx
	MT[215] = xxx
	MT[216] = xxx
	MT[217] = xxx
	MT[218] = xxx
	MT[219] = scn
	MT[220] = liv
	MT[221] = spl
	MT[222] = srt
	MT[223] = flr
	MT[224] = xxx
	MT[225] = xxx
	MT[226] = xxx
	MT[227] = xxx
	MT[228] = xxx
	MT[229] = xxx
	MT[230] = xxx
	MT[231] = xxx
	MT[232] = xxx
	MT[233] = xxx
	MT[234] = xxx
	MT[235] = kst
	MT[236] = lgf
	MT[237] = xxx
	MT[238] = xxx
	MT[239] = xxx
	MT[240] = prs
	MT[241] = xxx
	MT[242] = rnd
	MT[243] = xxx
	MT[244] = xxx
	MT[245] = xxx
	MT[246] = xxx
	MT[247] = xxx
	MT[248] = xxx
	MT[249] = xxx
	MT[250] = xxx
	MT[251] = ovr
	MT[252] = rev
	MT[253] = jon
	MT[254] = not
	MT[255] = xxx
	copy(MC[160:], " H@HHHHP@`HHHHHPDDDDDDDDDDH`HHHHHBBBBBBBBBBBBBBBBBBBBBBBBBB@P`HH@AAAAAAAAAAAAAAAAAAAAAAAAAA@H`H\x00")
}
