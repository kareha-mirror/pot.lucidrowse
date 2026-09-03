import 'package:flutter/material.dart';

import 'package:client/screens/create_screen.dart';
import 'package:client/screens/debug_screen.dart';
import 'package:client/screens/explore_screen.dart';
import 'package:client/screens/front_screen.dart';
import 'package:client/screens/home_screen.dart';
import 'package:client/screens/write_screen.dart';
import 'package:client/state.dart';

class App extends StatelessWidget {
  final AppState state;

  const App({super.key, required this.state});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ルシドロウズ - Lucidrowse',

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

      routes: {
        '/': (context) => FrontScreen(state: state),
        '/home': (context) => HomeScreen(state: state),

        '/create': (context) => CreateScreen(state: state),
        '/write': (context) => WriteScreen(state: state),
        '/explore': (context) => ExploreScreen(state: state),

        '/debug': (context) => DebugScreen(state: state),
      },
      initialRoute: '/',
    );
  }
}
