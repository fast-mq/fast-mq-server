// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

//                    ::
//                   :;J7, :,                        ::;7:
//                   ,ivYi, ,                       ;LLLFS:
//                   :iv7Yi                       :7ri;j5PL
//                  ,:ivYLvr                    ,ivrrirrY2X,
//                  :;r@Wwz.7r:                :ivu@kexianli.
//                 :iL7::,:::iiirii:ii;::::,,irvF7rvvLujL7ur
//                ri::,:,::i:iiiiiii:i:irrv177JX7rYXqZEkvv17
//             ;i:, , ::::iirrririi:i:::iiir2XXvii;L8OGJr71i
//           :,, ,,:   ,::ir@mingyi.irii:i:::j1jri7ZBOS7ivv,
//              ,::,    ::rv77iiiriii:iii:i::,rvLq@huhao.Li
//          ,,      ,, ,:ir7ir::,:::i;ir:::i:i::rSGGYri712:
//        :::  ,v7r:: ::rrv77:, ,, ,:i7rrii:::::, ir7ri7Lri
//       ,     2OBBOi,iiir;r::        ,irriiii::,, ,iv7Luur:
//     ,,     i78MBBi,:,:::,:,  :7FSL: ,iriii:::i::,,:rLqXv::
//     :      iuMMP: :,:::,:ii;2GY7OBB0viiii:i:iii:i:::iJqL;::
//    ,     ::::i   ,,,,, ::LuBBu BBBBBErii:i:i:i:i:i:i:r77ii
//   ,       :       , ,,:::rruBZ1MBBqi, :,,,:::,::::::iiriri:
//  ,               ,,,,::::i:  @arqiao.       ,:,, ,:::ii;i7:
// :,       rjujLYLi   ,,:::::,:::::::::,,   ,:i,:,,,,,::i:iii
// ::      BBBBBBBBB0,    ,,::: , ,:::::: ,      ,,,, ,,:::::::
// i,  ,  ,8BMMBBBBBBi     ,,:,,     ,,, , ,   , , , :,::ii::i::
// :      iZMOMOMBBM2::::::::::,,,,     ,,,,,,:,,,::::i:irr:i:::,
// i   ,,:;u0MBMOG1L:::i::::::  ,,,::,   ,,, ::::::i:i:iirii:i:i:
// :    ,iuUuuXUkFu7i:iii:i:::, :,:,: ::::::::i:i:::::iirr7iiri::
// :     :rk@Yizero.i:::::, ,:ii:::::::i:::::i::,::::iirrriiiri::,
//  :      5BMBBBBBBSr:,::rv2kuii:::iii::,:i:,, , ,,:,:i@petermu.,
//       , :r50EZ8MBBBBGOBBBZP7::::i::,:::::,: :,:,::i;rrririiii::
//           :jujYY7LS0ujJL7r::,::i::,::::::::::::::iirirrrrrrr:ii:
//        ,:  :@kevensun.:,:,,,::::i:i:::::,,::::::iir;ii;7v77;ii;i,
//        ,,,     ,,:,::::::i:iiiii:i::::,, ::::iiiir@xingjief.r;7:i,
//     , , ,,,:,,::::::::iiiiiiiiii:,:,:::::::::iiir;ri7vL77rrirri::
//      :,, , ::::::::i:::i:::i:i::,,,,,:,::i:i:::iir;@Secbone.ii:::

package cmd

import (
	"time"

	"github.com/fast-mq/server/internal/app"
	"github.com/higker/logker"
	"github.com/urfave/cli"
)

var (

	// whether enable  backgrund running
	isBackgrund bool

	// Start program start command.
	Start = cli.Command{
		// Command name
		Name: "start",
		// Command abbreviation
		Aliases: []string{"up"},
		// Command usage notes, here will show how to use the command when you enter the program name -help
		Usage: "FastMQ program start command✅.\n",
		// Command processing function
		Action: func(c *cli.Context) error {
			//c.Args().First()
			app.EchoInfo()

			if isBackgrund {
				logker.Info("%s", "FastMQ serve start up.")
				logker.Info("%s", "FastMQ is enable daemon mode running.")
			} else {
				logker.Info("%s", "FastMQ serve start up.")
			}
			time.Sleep(time.Second * 5)
			return nil
		},
	}
	// Stop program stop command.
	Stop = cli.Command{

		Name: "stop",

		Aliases: []string{"down"},

		Usage: "FastMQ program stop command❎.\n",

		Action: func(c *cli.Context) error {
			// STOP
			return nil
		},
	}
)

func init() {
	Start.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "daemon, d",
			Usage:       "Whether the program is running in the background.",
			Destination: &isBackgrund,
		},
	}
}
