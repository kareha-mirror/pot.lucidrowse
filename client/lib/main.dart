import 'package:flutter/material.dart';

enum InhabitStage { foreigner, inhabitant }

class AppState {
  InhabitStage stage = InhabitStage.foreigner;
  String profile = '';
  String text = '';
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
      theme: ThemeData(colorScheme: .fromSeed(seedColor: Colors.green)),
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
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/home'),
              child: Text('まどろみの水晶球'),
            ),
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/explore'),
              child: Text('光がおどる小箱'),
            ),
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/help'),
              child: Text('親切な鏡の精'),
            ),
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/debug'),
              child: Text('古びた道具箱'),
            ),
            Spacer(),
          ],
        ),
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
    if (widget.state.stage == InhabitStage.foreigner) {
      setState(() {
        widget.state.profile = value;
        widget.state.stage = InhabitStage.inhabitant;
        _controller.text = '';
      });
    } else {
      setState(() {
        widget.state.text = value;
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
        child: Column(
          mainAxisAlignment: .center,
          children: [
            Spacer(),
            if (widget.state.stage == InhabitStage.foreigner) ...[
              Text('水晶球が心に呼びかける。'),
              Text('夢の世界へいざないましょう。'),
              Text('この世界で夢見るあなたは、どんな人ですか、生きものですか。'),
            ] else ...[
              Text('水晶球が心に問いかける。'),
              Text('今日は何をしましたか？'),
            ],
            Spacer(),
            if (widget.state.stage == InhabitStage.foreigner)
              TextField(
                controller: _controller,
                decoration: InputDecoration(hintText: '森の薬草取りの少女。'),
                onSubmitted: (String value) => _submitText(value),
              )
            else
              TextField(
                controller: _controller,
                decoration: InputDecoration(hintText: '冒険者の泉で魚釣りをしてみた。'),
                onSubmitted: (String value) => _submitText(value),
              ),
            Spacer(),
            ElevatedButton(
              onPressed: () => _submitText(_controller.text),
              child: Text('答える'),
            ),
            if (widget.state.profile != '') ...[
              Spacer(),
              Text('${widget.state.profile}'),
            ],
            if (widget.state.text != '') ...[
              Spacer(),
              Text('${widget.state.text}'),
            ],
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('目をそらす'),
            ),
            Spacer(),
          ],
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
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('光がおどる小箱'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            Spacer(),
            Text('夢の世界の住人たち'),
            if (widget.state.stage == InhabitStage.inhabitant) ...[
              Spacer(),
              Text('${widget.state.profile}'),
              Text('${widget.state.text}'),
            ] else ...[
              Spacer(),
              Text('まだ誰も住んでない。'),
            ],
            Spacer(),
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
            Spacer(),
            Text('まどろみの水晶球をのぞき込むと、夢の世界につながるよ。'),
            Text('そこでは、みんなの夢がひとつになるんだ。'),
            Spacer(),
            Text('光がおどる小箱を開けると、この世界に住むに仲間たちに出会えるよ。'),
            Text('話しかけることはできないけど、すれちがうことはあるかもしれない。'),
            Spacer(),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('目をそらす'),
            ),
            Spacer(),
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
        title: Text('古びた道具箱'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            Spacer(),
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
