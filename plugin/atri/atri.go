/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"encoding/base64"
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
)

type datagetter func(string, bool) ([]byte, error)

func (dgtr datagetter) randImage(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.ImageBytes(data)
}

func (dgtr datagetter) randRecord(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.Record("base64://" + base64.StdEncoding.EncodeToString(data))
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

// isAtriSleeping 凌晨0点到6点，ATRI 在睡觉，不回应任何请求
func isAtriSleeping(*zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 1 && now < 6 {
		return false
	}
	return true
}

func init() { // 插件主体
	engine := control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "atri人格文本回复",
		Help: "本插件基于 ATRI ，为 Golang 移植版\n" +
			"- ATRI醒醒\n- ATRI睡吧\n- 萝卜子\n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你mua~ | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爬 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
		PublicDataFolder: "Atri",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("嗯呜呜……先生……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	engine.UsePreHandler(isAtriSleeping)
	var dgtr datagetter = engine.GetLazyData
	engine.OnFullMatch("小矮子").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(randText("小矮子是对机器人的蔑称！", "是瑶瑶......小矮子可是对机器人的蔑称"))
			case 1:
				ctx.SendChain(dgtr.randRecord("RocketPunch.amr"))
			}
		})
	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き", "贴贴", "老婆", "亲一个", "mua~"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("SUKI.jpg", "SUKI1.jpg", "SUKI2.png"))
		})
	engine.OnKeywordGroup([]string{"草你妈", "操你马听见了吗", "脑瘫", "废柴", "fw", "five", "废物", "煞笔", "爬", "爪巴", "sb", "SB", "傻B"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("FN.jpg", "WQ.jpg", "WQ1.jpg"))
		})
	engine.OnFullMatchGroup([]string{"早安", "早！又是元气满满的一天呢~", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早哦", "早早早"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...不要了啦..zzz....",
					"别...不要..zzz..那..zzz..不可以....",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
					"唔...好可恶..把人家弄醒了啦！",
					"大半夜找我有何贵干！杂~鱼~（超生气的说）",
				))
			case now >= 6 && now < 9:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...(哈欠)",
					"唔......吧唧...早上...哈啊啊~~~\n早上好......",
					"早上好......",
					"早上好呜......呼啊啊~~~~",
					"啊......早上好。\n昨晚也很激情呢！",
					"吧唧吧唧......怎么了...已经早上了么...",
					"早上好！",
					"......看起来像是傍晚，其实已经早上了吗？",
					"早上好......欸~~~脸好近呢",
					"早~又是元气满满的一天，今天也要加油哦！",
					"唔~~怎么到早上了啦！都怪你昨晚太费时了啦~（哼！生气ing）",
				))
			case now >= 9 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"哼！这个点还早啥，昨晚干啥去了！？",
					"熬夜了对吧熬夜了对吧熬夜了对吧？？？！",
					"是不是熬夜了昨晚一定很棒吧❥(^_-)",
					"早什么早！要说你昨晚好棒！",
					"昨晚是不是对着我偷偷的爱爱了？！\n哼！一定是！居然这么晚才起床！",
					"怎么现在才起床？昨晚对我说“晚安”都是假的吗？！\n终究是错付了，呜呜呜呜（哭的好大声）",
					"这个点对我说早，你个小懒猪！你的心里根本就没有我~\n哼！超生气！哄不好的那种！！",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"早个啥？哼唧！我都准备洗洗睡了！",
					"不是...你看看几点了，哼！",
					"晚上好哇",
					"这个点你的女神都睡了，你居然还在熬夜！哦~你没有女朋友对吧！怪不得没有女朋友！！",
					"晚安捏，梦里希望有我~诶嘿~(*╹▽╹*)",
					"呜呜呜，有一天结束了捏，时间过得好快，不知不觉又老了一天~",
					"晚安，要早睡哦，明天依旧是光芒万丈捏嘻嘻(#^.^#)",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，我会一直陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
					"去睡你的午觉吧！我才不需要你陪呢！才！不！需要！！",
					"午安，好孩子要坚持午休哦，只有这样小脑袋瓜子才可以保持清醒呢。",
					"如果说爱是一种无形的等待，那么我愿意一直等着你，直到你能够想起我。",
					"午安，我会一直等着你，就像。。。就像你喜欢和我聊天一样(*^▽^*)~",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好讨厌...让我睡觉好不好~求求了~..zzz....",
					"别...不要..zzz..不...不可以~..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....别闹了啦....人家要睡觉了啦~",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"你可猝死算了吧！",
					"？啊这",
					"亲，这边建议赶快去睡觉呢~~~",
					"可恶居然还不睡觉！！为何这还不猝死啊！！",
					"可恶的咸鱼现在居然还不睡，心里是放不下远方吗？\n快睡吧~只有保持清醒的头脑才可以向往远方哦，答应我待来日方长再来看我可好？晚安遇见你真好~",
					"亲，再不睡太阳公公可就下班了哦~",
					"这时候还不睡的，是不要命的肥宅吗？真是一群杂男鱼❥",
					"",
					"",
				))
			case now >= 11 && now < 15:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，ATRI会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
					"去睡你的午觉吧！我才不需要你陪呢！才！不！需要！！",
					"午安，好孩子要坚持午休哦，只有这样小脑袋瓜子才可以保持清醒呢。",
					"如果说爱是一种无形的等待，那么我愿意一直等着你，直到你能够想起我。",
					"午安，我会一直等着你，就像。。。就像你喜欢和我聊天一样(*^▽^*)~",
				))
			case now >= 15 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"难不成？？晚上不想睡觉？？现在休息",
					"就......挺离谱的...现在睡觉",
					"现在还是白天哦，睡觉还太早了",
					"不！可！以！睡！起！来！陪！我！玩！",
					"现在是白天哦，做那种事情会不会太早了呀。。。",
					"",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"嗯？...嗯~....(打瞌睡)",
					"呼...呼...已经睡着了哦~...呼......",
					"......我、我会在这守着你的，请务必好好睡着",
					"晚安捏，梦里希望有我~诶嘿~(*╹▽╹*)",
					"呜呜呜，有一天结束了捏，时间过得好快，不知不觉又老了一天~",
					"晚安，要早睡哦，明天依旧是光芒万丈捏嘻嘻(#^.^#)",
				))
			}
		})
	engine.OnKeywordGroup([]string{"高性能", "太棒了", "すごい", "sugoi", "斯国一", "よかった"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"小事一桩，我是高性能的嘛",
				"怎么样？还是我比较高性能吧？",
				"哼哼！我果然是高性能的呢！",
				"因为我是高性能的嘛！嗯哼！",
				"因为我是高性能的呢！",
				"就是说啊~本宝宝最棒啦！",
				"哎呀~，我可真是太高性能了",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"嘿嘿嘿，璐瑶姐姐也是这么说的捏~",
				"嘿嘿，我的高性能发挥出来啦♪",
				"我果然是很高性能的机器人吧！",
				"是吧！谁叫我这么高性能呢！哼哼！",
				"交给我吧，有高性能的我陪着呢",
				"呣......我的高性能，毫无遗憾地施展出来了......",
			))
		})
	engine.OnKeywordGroup([]string{"没事", "没关系", "大丈夫", "还好", "不要紧", "没出大问题", "没伤到哪"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"没事没事，因为我是高性能的嘛！嗯哼！",
				"没事的，因为我是高性能的呢！",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"那种程度的事不算什么的。\n别看我这样，我可是高性能的",
				"没问题的，我可是高性能的",
			))
		})

	engine.OnKeywordGroup([]string{"好吗", "是吗", "行不行", "能不能", "可不可以"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("YES.png", "NO.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"啊这"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("AZ.jpg", "AZ1.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"我好了"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText("不许好！", "憋回去！"))
		})
	engine.OnFullMatchGroup([]string{"？", "?", "¿"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(dgtr.randImage("WH.jpg", "WH1.jpg", "WH2.jpg", "WH3.jpg"))
			}
		})
	engine.OnKeyword("离谱").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(dgtr.randImage("WH.jpg"))
			}
		})
	engine.OnKeyword("答应我", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText("我无法回应你的请求"))
		})
}
