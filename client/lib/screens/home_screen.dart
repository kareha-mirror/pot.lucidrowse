import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';

class HomeScreen extends StatefulWidget {
  final AppState state;

  const HomeScreen({super.key, required this.state});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
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
              child: const Image(image: AssetImage('assets/images/home.webp')),
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
                    onPressed: () => Navigator.pushNamed(context, '/create'),
                    child: Text('まどろみの水晶球'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/write'),
                    child: Text('やわらかい羽ペン'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/read'),
                    child: Text('おしゃれな日記帳'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/explore'),
                    child: Text('光おどる小箱'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) =>
                            const HelpScreen(page: HelpPage.home),
                      ),
                    ),
                    child: Text('親切な鏡の精'),
                  ),
                  SizedBox(height: 96),
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: Text('外を見る'),
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
