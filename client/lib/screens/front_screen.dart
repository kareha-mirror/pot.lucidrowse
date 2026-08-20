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
        title: const Text('ルシドロウズ'),
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
              child: const Padding(
                padding: EdgeInsets.symmetric(horizontal: 24),
                child: Column(
                  children: [
                    SizedBox(height: 24),
                    Image(image: AssetImage('assets/images/logo.webp')),
                  ],
                ),
              ),
            ),
          ),

          Center(
            child: Column(
              mainAxisAlignment: .center,
              children: [
                ElevatedButton(
                  onPressed: () => Navigator.pushNamed(context, '/home'),
                  child: const Text('古びた宿'),
                ),

                const SizedBox(height: 48),

                helpButton(context, HelpPage.front),
              ],
            ),
          ),

          Positioned(
            left: 0,
            right: 0,
            bottom: MediaQuery.paddingOf(context).bottom + 24,
            child: Center(
              child: ElevatedButton(
                onPressed: () => Navigator.pushNamed(context, '/debug'),
                child: const Text('汚れた道具箱'),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
