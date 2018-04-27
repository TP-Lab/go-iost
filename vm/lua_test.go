package vm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/iost-official/prototype/state/mocks"
	"github.com/golang/mock/gomock"
	"github.com/iost-official/prototype/state"
	"github.com/iost-official/gopher-lua"
)

func TestLuaVM(t *testing.T) {
	Convey("Test of Lua VM", t, func() {
		mockCtl := gomock.NewController(t)
		pool := state_mock.NewMockPool(mockCtl)

		var k state.Key
		var v state.Value

		pool.EXPECT().Put(gomock.Any(), gomock.Any()).AnyTimes().Do(func(key state.Key, value state.Value) error {
			k = key
			v = value
			return nil
		})
		pool.EXPECT().Copy().AnyTimes().Do(func() state.Pool {
			return pool
		})
		main := NewLuaMethod("main", 1)
		lc := LuaContract{
			info: ContractInfo{},
			code:
			`function main()
	Put("hello", "world")
	return "success"
end`,
			main: &main,
		}

		lvm := LuaVM{}
		lvm.Prepare(&lc, pool, "test")
		lvm.Pool = pool
		lvm.Start()
		ret, _, err := lvm.Call("main")
		lvm.Stop()

		So(err, ShouldBeNil)
		So(ret[0].String(), ShouldEqual, "success")
		So(k, ShouldEqual, "testhello")
		So(v.String(), ShouldEqual, "world")

	})

	Convey("Test of Lua value converter", t, func() {
		Convey("Lua to core", func() {
			Convey("string", func() {
				lstr := lua.LString("hello")
				cstr := Lua2Core(lstr)
				So(cstr.Type(), ShouldEqual, state.String)
				So(cstr.String(), ShouldEqual, "hello")
			})
		})
	})
}
