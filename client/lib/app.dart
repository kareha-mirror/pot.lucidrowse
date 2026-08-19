import 'package:flutter/material.dart';

import 'package:client/screens/create_screen.dart';
import 'package:client/screens/debug_screen.dart';
import 'package:client/screens/explore_screen.dart';
import 'package:client/screens/front_screen.dart';
//import 'package:client/screens/help_screen.dart';
import 'package:client/screens/home_screen.dart';
import 'package:client/screens/read_screen.dart';
import 'package:client/screens/write_screen.dart';
import 'package:client/state/app_state.dart';

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
        '/': (context) => const FrontScreen(),
        '/home': (context) => HomeScreen(state: state),

        '/create': (context) => CreateScreen(state: state),
        '/write': (context) => WriteScreen(state: state),
        '/read': (context) => ReadScreen(state: state),
        '/explore': (context) => ExploreScreen(state: state),

        //'/help': (context) => const HelpScreen(),
        '/debug': (context) => DebugScreen(state: state),
      },
    );
  }
}
