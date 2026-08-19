import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';

class FrontScreen extends StatefulWidget {
  const FrontScreen({super.key});

  @override
  State<FrontScreen> createState() => _FrontScreenState();
}

class _FrontScreenState extends State<FrontScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('ルシドロウズ'),
      ),
      body: Stack(
        children: [
          Center(
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 1200),
              child: const Image(image: AssetImage('assets/images/front.webp')),
            ),
          ),
          Center(
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 400),
              child: Padding(
                padding: EdgeInsets.symmetric(horizontal: 24),
                child: Column(
                  mainAxisAlignment: .center,
                  children: [
                    SizedBox(height: 24),
                    const Image(image: AssetImage('assets/images/logo.webp')),
                    Spacer(),
                  ],
                ),
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
                    child: Text('古びた宿'),
                  ),
                  SizedBox(height: 48),
                  ElevatedButton(
                    onPressed: () => Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) =>
                            const HelpScreen(page: HelpPage.front),
                      ),
                    ),
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
