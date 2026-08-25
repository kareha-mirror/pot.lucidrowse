import 'package:flutter/material.dart';

import 'package:client/app.dart';
import 'package:client/state.dart';

void main() {
  final state = AppState();
  runApp(App(state: state));
}
