import 'package:flutter/material.dart';

enum InhabitKind {
  foreigner,
  inhabitant,
  hermit,
  forgotten;

  bool get isForeigner => this == InhabitKind.foreigner;
}

class Player {
  InhabitKind inhabitKind;
  String flavorText;
  String actionText;

  Player.unnamed()
    : inhabitKind = InhabitKind.foreigner,
      flavorText = '',
      actionText = '';
}

class AppState {
  Player player = Player.unnamed();
  List<Player> inhabitants = [];
}

void main() {
  final state = AppState();
  runApp(MyApp(state: state));
}

class MyApp extends StatelessWidget {
  final AppState state;

  const MyApp({super.key, required this.state});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ルシドロウズ',
      theme: ThemeData(
        colorScheme: .fromSeed(
          seedColor: Colors.purple,
          brightness: Brightness.light,
        ),
      ),
      darkTheme: ThemeData(
        colorScheme: .fromSeed(
          seedColor: Colors.purple,
          brightness: Brightness.dark,
        ),
      ),
      themeMode: ThemeMode.system,
      initialRoute: '/',
      routes: {
        '/': (context) => const FrontPage(),
        '/home': (context) => HomePage(state: state),
        '/explore': (context) => ExplorePage(state: state),
        '/help': (context) => const HelpPage(),
        '/debug': (context) => DebugPage(state: state),
      },
    );
  }
}

class FrontPage extends StatefulWidget {
  const FrontPage({super.key});

  @override
  State<FrontPage> createState() => _FrontPageState();
}

class _FrontPageState extends State<FrontPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('古びた宿'),
      ),
      body: Stack(
        children: [
          Center(
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 1200),
              child: Column(
                mainAxisAlignment: .center,
                children: [
                  const Image(image: AssetImage('assets/images/front.webp')),
                  Spacer(),
                ],
              ),
            ),
          ),
          Center(
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 1200),
              child: Column(
                mainAxisAlignment: .center,
                children: [
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/home'),
                    child: Text('まどろみの水晶球'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/explore'),
                    child: Text('光おどる小箱'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/help'),
                    child: Text('親切な鏡の精'),
                  ),
                  SizedBox(height: 96),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/debug'),
                    child: Text('汚れた道具箱'),
                  ),
                  SizedBox(height: 48),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class HomePage extends StatefulWidget {
  final AppState state;

  const HomePage({super.key, required this.state});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _submitText(String value) {
    if (widget.state.player.inhabitKind.isForeigner) {
      setState(() {
        widget.state.player.flavorText = value;
        widget.state.player.inhabitKind = InhabitKind.inhabitant;
        _controller.text = '';

        widget.state.inhabitants.add(widget.state.player);
      });
    } else {
      setState(() {
        widget.state.player.actionText = value;
        _controller.text = '';
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('まどろみの水晶球'),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              SizedBox(height: 48),
              if (widget.state.player.inhabitKind.isForeigner) ...[
                Text('水晶球が心に呼びかける。'),
                Text('夢の世界へいざないましょう。'),
                SizedBox(height: 24),
                Text('この世界で夢見るあなたは、どんな人ですか、どんな生きものですか。'),
              ] else ...[
                Text('水晶球が心に問いかける。'),
                SizedBox(height: 24),
                Text('今日は何をして過ごしましたか？'),
              ],
              SizedBox(height: 24),
              if (widget.state.player.inhabitKind.isForeigner)
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 48),
                  child: TextField(
                    controller: _controller,
                    decoration: InputDecoration(hintText: '森の薬草取りの少女。'),
                    onSubmitted: (String value) => _submitText(value),
                    onChanged: (String value) => setState(() {}),
                    maxLines: null,
                  ),
                )
              else
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 48),
                  child: TextField(
                    controller: _controller,
                    decoration: InputDecoration(hintText: '冒険者の泉で魚釣りをしてみた。'),
                    onSubmitted: (String value) => _submitText(value),
                    onChanged: (String value) => setState(() {}),
                    maxLines: null,
                  ),
                ),
              SizedBox(height: 24),
              if (_controller.text == '')
                ElevatedButton(onPressed: null, child: Text('答える'))
              else
                ElevatedButton(
                  onPressed: () => _submitText(_controller.text),
                  child: Text('答える'),
                ),
              if (widget.state.player.flavorText != '') ...[
                SizedBox(height: 48),
                Text('${widget.state.player.flavorText}'),
              ],
              if (widget.state.player.actionText != '') ...[
                SizedBox(height: 48),
                Text('${widget.state.player.actionText}'),
              ],
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pushNamed(context, '/'),
                child: Text('目をそらす'),
              ),
              SizedBox(height: 48),
            ],
          ),
        ),
      ),
    );
  }
}

class ExplorePage extends StatefulWidget {
  final AppState state;

  const ExplorePage({super.key, required this.state});

  @override
  State<ExplorePage> createState() => _ExplorePageState();
}

class _ExplorePageState extends State<ExplorePage> {
  List<Widget> inhabitantsList() {
    List<Widget> list = [];
    for (final inhabitant in widget.state.inhabitants) {
      list.add(SizedBox(height: 24));
      list.add(Text('${inhabitant.flavorText}'));
      list.add(Text('${inhabitant.actionText}'));
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('光おどる小箱'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            SizedBox(height: 48),
            Text('夢の世界の住人たち'),
            if (widget.state.inhabitants.isEmpty) ...[
              SizedBox(height: 48),
              Text('まだ誰も住んでない。'),
            ] else ...[
              SizedBox(height: 24),
              ...inhabitantsList(),
            ],
            SizedBox(height: 96),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('ふたを閉じる'),
            ),
            SizedBox(height: 48),
          ],
        ),
      ),
    );
  }
}

class HelpPage extends StatefulWidget {
  const HelpPage({super.key});

  @override
  State<HelpPage> createState() => _HelpPageState();
}

class _HelpPageState extends State<HelpPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('親切な鏡の精'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            SizedBox(height: 48),
            Text('まどろみの水晶球をのぞき込むと、夢の世界につながるよ。'),
            Text('そこでは、みんなの夢がひとつになるんだ。'),
            SizedBox(height: 48),
            Text('光おどる小箱を開けると、この世界に住む仲間たちに出会えるよ。'),
            Text('話しかけることはできないけど、すれちがうことはあるかもしれないね。'),
            SizedBox(height: 96),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('目をそらす'),
            ),
            SizedBox(height: 48),
          ],
        ),
      ),
    );
  }
}

class DebugPage extends StatefulWidget {
  final AppState state;

  const DebugPage({super.key, required this.state});

  @override
  State<DebugPage> createState() => _DebugPageState();
}

class _DebugPageState extends State<DebugPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('汚れた道具箱'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            SizedBox(height: 48),
            if (widget.state.player.inhabitKind.isForeigner)
              ElevatedButton(onPressed: null, child: Text('自分を手放す'))
            else
              ElevatedButton(
                onPressed: () =>
                    setState(() => widget.state.player = Player.unnamed()),
                child: Text('自分を手放す'),
              ),
            SizedBox(height: 48),
            if (widget.state.inhabitants.isEmpty)
              ElevatedButton(onPressed: null, child: Text('住人たちを追い出す'))
            else
              ElevatedButton(
                onPressed: () => setState(() {
                  widget.state.player = Player.unnamed();
                  widget.state.inhabitants = [];
                }),
                child: Text('住人たちを追い出す'),
              ),
            SizedBox(height: 96),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('ふたを閉じる'),
            ),
            Spacer(),
          ],
        ),
      ),
    );
  }
}
