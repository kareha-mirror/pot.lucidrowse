import 'package:flutter/material.dart';

enum HelpPage { front, home }

class HelpScreen extends StatefulWidget {
  const HelpScreen({super.key, required this.page});

  final HelpPage page;

  @override
  State<HelpScreen> createState() => _HelpScreenState();
}

class _HelpScreenState extends State<HelpScreen> {
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
            if (widget.page == HelpPage.home) ...[
              SizedBox(height: 48),
              Text('まどろみの水晶球をのぞき込むと、夢の世界につながるよ。'),
              Text('そこでは、みんなの夢がひとつになるんだ。'),
              SizedBox(height: 48),
              Text('光おどる小箱を開けると、この世界に住む仲間たちに出会えるよ。'),
              Text('話しかけることはできないけど、すれちがうことはあるかもしれないね。'),
            ],
            SizedBox(height: 96),
            ElevatedButton(
              onPressed: () => Navigator.pop(context),
              child: Text('鏡を閉じる'),
            ),
            SizedBox(height: 48),
          ],
        ),
      ),
    );
  }
}
