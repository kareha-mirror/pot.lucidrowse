import 'package:flutter/material.dart';

enum HelpPage { front, home, debug }

class HelpScreen extends StatefulWidget {
  const HelpScreen({super.key, required this.page});

  final HelpPage page;

  @override
  State<HelpScreen> createState() => _HelpScreenState();
}

List<Widget> _helpFront = [
  const Text('ようこそ、ルシドロウズへ。'),
  const Text('ここは夢と現実の境界だよ。'),
  const Text('夢の世界の住人になって、世界と自由に関われるんだ。'),
  const SizedBox(height: 48),
  const Text('古びた宿に泊まれば夢の世界の住人になれるよ。'),
  const SizedBox(height: 24),
  const Text('この世界に住むには、まず夢の世界と関わるための分身を作るんだ。'),
  const Text('そして毎日一度だけ、その日に何をしたかを日記に書く。'),
  const Text('書いたことがそのまま夢の現実になるとは限らないけど、ある程度は自由だよ。'),
  const Text('次の日目を覚ましたときに何が本当の夢になったか確認できる。'),
  const SizedBox(height: 24),
  const Text('あなたの他にもここには色んな住人たちが住んでいるよ。'),
  const Text('近所の住人たちのことを眺めることもできる。'),
  const Text('他の住人たちに直接話しかけることはできないけど、毎日の行動の中では関わることもあるかもね。'),
  const SizedBox(height: 48),
  const Text('汚れた道具箱にはこの世界の元の管理人が使っていた魔法の道具たちが入ってる。'),
  const Text('これを使うと色んなずるいことができるけど、夢の世界を壊してしまうこともあるから注意して。'),
  const Text('これはたぶんそのうち封印されるだろうね。'),
];

List<Widget> _helpHome = [
  const Text('まどろみの水晶球をのぞき込むと、夢の世界につながるよ。'),
  const Text('そこでは、みんなの夢がひとつになるんだ。'),
  const SizedBox(height: 48),
  const Text('光おどる小箱を開けると、この世界に住む仲間たちに出会えるよ。'),
  const Text('話しかけることはできないけど、すれちがうことはあるかもしれないね。'),
];

List<Widget> _helpDebug = [
  const Text('「自分を手放す」で夢の世界の分身を解放すれば、分身を作り直せるよ。'),
  const Text('それでも解放された分身は夢の世界の住人として残るよ。'),
  const SizedBox(height: 48),
  const Text('「住人たちを追い出す」で夢の世界の住人たちを空っぽに戻せるよ。'),
];

class _HelpScreenState extends State<HelpScreen> {
  List<Widget> _helpWidgets(HelpPage page) {
    switch (page) {
      case HelpPage.front:
        return _helpFront;
      case HelpPage.home:
        return _helpHome;
      case HelpPage.debug:
        return _helpDebug;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('親切な鏡の精'),
      ),

      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              ..._helpWidgets(widget.page),

              const SizedBox(height: 96),

              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: const Text('鏡を閉じる'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
