import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';

class Player {
  final String id;
  final int day;
  Flavor? flavor;
  PlayerAction? action;

  Player({required this.id, required this.day});
}
