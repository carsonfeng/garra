package a

type IntSet map[int]struct{}

type genProtoUserConfig struct {
	GetRoomIDFromCache bool
	MyBanSet           IntSet
	MyFollowSet        IntSet
	MyLikeSet          IntSet
	MyStarFollowSet    IntSet
}

var (
	defaultGenProtoUserConfig = genProtoUserConfig{
		GetRoomIDFromCache: true,
	}
)

func init() {
	defaultGenProtoUserConfig.MyBanSet = map[int]struct{}{}
	defaultGenProtoUserConfig.MyStarFollowSet = map[int]struct{}{}
}

func ban0() {
	defaultGenProtoUserConfig.MyBanSet = map[int]struct{}{}

	defaultGenProtoUserConfig.MyStarFollowSet = map[int]struct{}{}
}

func pass1() {
	cfg := defaultGenProtoUserConfig
	cfg.MyBanSet = map[int]struct{}{}
}

func pass2() {
	cfg := defaultGenProtoUserConfig
	cfg2 := cfg
	cfg2.MyBanSet = map[int]struct{}{}
}

func ban1() {
	cfg := &defaultGenProtoUserConfig
	cfg.MyBanSet = map[int]struct{}{}
}

func ban2() {
	cfg := &defaultGenProtoUserConfig
	cfg2 := cfg
	cfg2.MyBanSet = map[int]struct{}{}
}

func ban3() {
	cfg := &defaultGenProtoUserConfig
	cfg2 := cfg
	var cfg3 *genProtoUserConfig
	cfg3 = cfg2
	cfg3.MyBanSet = map[int]struct{}{}
}

//func modifyFunc(cfg *genProtoUserConfig) {
//	cfg.MyLikeSet = map[int]struct{}{}
//}
//
//func ban3() {
//	cfg := defaultGenProtoUserConfig
//	modifyFunc(&cfg)
//}
