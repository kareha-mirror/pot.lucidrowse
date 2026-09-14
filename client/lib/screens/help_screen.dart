import 'package:flutter/material.dart';

import 'package:client/widgets/translucent_panel.dart';

enum HelpPage { front, home, key, debug }

class HelpScreen extends StatefulWidget {
  const HelpScreen({super.key, required this.page});

  final HelpPage page;

  @override
  State<HelpScreen> createState() => _HelpScreenState();
}

List<Widget> _helpFront = [
  TranslucentPanel(
    child: Column(
      children: [
        const Text('ようこそ、ルシドロウズへ。'),
        const Text('ここは夢と現実の境界。'),
        const Text('夢の世界の住人になって、世界と自由に関わることができるんだ。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(child: const Text('古びた宿に泊まれば夢の世界の住人になれるよ。')),

  const SizedBox(height: 24),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('この世界に住むには、まず夢の世界と関わるための分身を作るんだ。'),
        const Text('そして毎日一度だけ、その日に何をしたかを日記に書く。'),
        const Text('するとそれが夢の真実になる。'),
        const Text('書いたことがそのまま夢の現実になるとは限らないけれど、ある程度は自由に振る舞える。'),
        const Text('次の日、目を覚ましたときに、みんなの夢が共有されるから、楽しみに待とうね。'),
      ],
    ),
  ),

  const SizedBox(height: 24),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('あなたの他にもここには色んな住人たちが住んでいる。'),
        const Text('近所の住人たちのことを眺めることもできるから、気になるなら様子を見てみたらどうかな。'),
        const Text('他の住人たちに直接話しかけることはできないけれど、毎日の行動の中では関わることもあるかもしれない。'),
        const Text('緩く構えて行こう。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('汚れた道具箱には夢の魔法使いが使っていた秘密の道具たちが入っている。'),
        const Text('これを使うと色々なずるができるけれど、夢の世界を壊してしまうこともあるから注意してほしい。'),
        const Text('これはたぶんそのうち封印されるだろうね。'),
      ],
    ),
  ),
];

List<Widget> _helpHome = [
  TranslucentPanel(
    child: Column(
      children: [
        const Text('まどろみの水晶球をのぞき込むと、夢の世界につながるんだ。'),
        const Text('そこでは、みんなの夢がひとつになる。'),
        const Text('どんな存在として夢と関わりたいか念じれば、それがあなたの分身の姿になる。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('やわらかい羽ペンで日記を書くことができる。'),
        const Text('次の日までの間に夢の世界が日記を読んで、それを真実に変える。'),
        const Text('もちろん、すべてがそのままではないけどね。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(children: [const Text('おしゃれな日記帳を開けば、真実になった日記を読むことができる。')]),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('光おどる小箱を開けると、この世界に住む仲間たちに出会える。'),
        const Text('話しかけることはできないけれど、すれちがうことはあるかもしれないね。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('1日の日記の中で移動できるのは、1日で往復できる範囲だけだから、地域を越えて移動することはできないよ。'),
        const Text('数日間日記を書くと、もう一度水晶球に念じることができる。'),
        const Text('そこで別の地域に引っ越すこともできるよ。'),
      ],
    ),
  ),
];

List<Widget> _helpKey = [
  TranslucentPanel(
    child: Column(
      children: [
        const Text('「秘密の鍵」を作っておけば、別の入口からでも同じ夢に入れるよ。'),
        const Text('秘密の鍵が無ければ、別の入口から入ると別な夢になってしまうし、1年経つと入れなくなってしまうんだ。'),
        const Text('今見ている夢とずっと一緒にいたいなら、秘密の鍵は作っておくと良いね。'),
      ],
    ),
  ),
];

List<Widget> _helpDebug = [
  TranslucentPanel(
    child: Column(
      children: [
        const Text('「自分を手放す」で夢の世界の分身を解放すれば、分身を作り直せる。'),
        const Text('それでも解放された分身は夢の世界の住人として残るんだ。'),
        const Text('解放された分身には他の人が入り込んで、また夢を見ることもあるよ。'),
      ],
    ),
  ),

  const SizedBox(height: 48),

  TranslucentPanel(
    child: Column(
      children: [
        const Text('「夢の世界に呼びかける」と夢が答えてくれる。'),
        const Text('答えてくれないときは夢を見れないよ。'),
      ],
    ),
  ),
];

class _HelpScreenState extends State<HelpScreen> {
  List<Widget> _helpWidgets(HelpPage page) {
    switch (page) {
      case .front:
        return _helpFront;
      case .home:
        return _helpHome;
      case .key:
        return _helpKey;
      case .debug:
        return _helpDebug;
    }
  }

  @override
  Widget build(BuildContext context) {
    return PopScope(
      canPop: false,
      onPopInvokedWithResult: (didPop, result) {
        if (!didPop) {
          Navigator.pop(context);
        }
      },
      child: Scaffold(
        body: Stack(
          children: [
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: const Image(
                image: AssetImage('assets/images/help.webp'),
                fit: .cover,
              ),
            ),

            SingleChildScrollView(
              padding: EdgeInsets.all(
                MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
              ),
              child: Center(
                child: Column(
                  mainAxisAlignment: .center,
                  children: [
                    ..._helpWidgets(widget.page),

                    const SizedBox(height: 96),

                    ElevatedButton(
                      onPressed: () => Navigator.pop(context),
                      child: const Text('納得した'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

Widget helpButton(BuildContext context, HelpPage page) {
  return ElevatedButton(
    onPressed: () => Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => HelpScreen(page: page)),
    ),
    child: const Text('親切な鏡の精'),
  );
}
