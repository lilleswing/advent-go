package main

import "testing"

func TestSet(t *testing.T) {
	regs := initializeRegisters()
	set(regs, "a", 28)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	if regs["a"] != 28 {
		t.Errorf("Must Set Value")
	}
}

func TestAdd(t *testing.T) {
	regs := initializeRegisters()
	add(regs, "a", 28)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	if regs["a"] != 28 {
		t.Error()
	}
}

func TestMul(t *testing.T) {
	regs := initializeRegisters()
	regs["p"] = 1
	mul(regs, "p", 28)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	if regs["p"] != 28 {
		t.Error()
	}
}

func TestMod(t *testing.T) {
	regs := initializeRegisters()
	regs["p"] = 28
	mod(regs, "p", 5)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	if regs["p"] != 3 {
		t.Error()
	}
}

func TestSend(t *testing.T) {
	regs := initializeRegisters()
	snd(regs, 5)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	if regs["snd"] != 5 {
		t.Error()
	}
}

func TestRcvEmpty(t *testing.T) {
	regs := initializeRegisters()
	rcv(regs, "b")
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	_, ok := regs["rcv"]
	if ok {
		t.Errorf("Should be noop")
	}
}

func TestRcvFull(t *testing.T) {
	regs := initializeRegisters()
	regs["snd"] = 5
	regs["b"] = 1
	rcv(regs, "b")
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
	v, ok := regs["rcv"]
	if !ok {
		t.Errorf("Should not be noop")
	}
	if v != 5 {
		t.Errorf("should be value in snd")
	}
}

func TestJgzZero(t *testing.T) {
	regs := initializeRegisters()
	jgz(regs, "b", -1)
	if regs["index"] != 1 {
		t.Errorf("Must increment index")
	}
}

func TestJgzGZero(t *testing.T) {
	regs := initializeRegisters()
	regs["b"] = 1
	jgz(regs, "b", -1)
	if regs["index"] != -1 {
		t.Errorf("wrong index")
	}
}

func TestSnd2Full(t *testing.T) {
	regs := initializeRegisters()
	regs["cnt"] = 0
	messages := make(chan int, 1)
	messages <- 1
	snd2(regs, 2, messages)
	if regs["index"] != 0 {
		t.Errorf("wrong index")
	}
	if regs["cnt"] != 0 {
		t.Errorf("wrong index")
	}
}

func TestSnd2Success(t *testing.T) {
	regs := initializeRegisters()
	regs["cnt"] = 0
	messages := make(chan int, 1)
	snd2(regs, 2, messages)
	if regs["index"] != 1 {
		t.Errorf("wrong index")
	}
	if regs["cnt"] != 1 {
		t.Errorf("wrong index")
	}
	v := <-messages
	if v != 2 {
		t.Errorf("wrong value on channel")
	}
}

func TestRcv2Fail(t *testing.T) {
	regs := initializeRegisters()
	messages := make(chan int, 1)
	rcv2(regs, "b", messages)
	if regs["index"] != 0 {
		t.Errorf("wrong index")
	}
	if regs["b"] != 0 {
		t.Errorf("wrong index")
	}
}

func TestRcv2Success(t *testing.T) {
	regs := initializeRegisters()
	messages := make(chan int, 1)
	messages <- 42
	rcv2(regs, "b", messages)
	if regs["index"] != 1 {
		t.Errorf("wrong index")
	}
	if regs["b"] != 42 {
		t.Errorf("wrong index")
	}
}

func TestRcv2SuccessMulti(t *testing.T) {
	regs := initializeRegisters()
	messages := make(chan int, 2)
	messages <- 42
	messages <- 41
	rcv2(regs, "b", messages)
	if regs["index"] != 1 {
		t.Errorf("wrong index")
	}
	if regs["b"] != 42 {
		t.Errorf("wrong index")
	}
}
