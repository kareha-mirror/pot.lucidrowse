import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';

Future<Player?> apiLoadPlayer() async {
  final response = await http.get(Uri.parse('$apiBase/players'));

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  final result = jsonDecode(response.body);

  if (result['id'] == null) {
    return null;
  }

  final player = Player(id: result['id'], day: result['day']);

  if (result['flavor'] != null) {
    final flavor = Flavor.fromJson(result['flavor']);
    flavor.imageId = result['flavor-image-id'];
    player.flavor = flavor;
  }

  if (result['action'] != null) {
    final action = PlayerAction();
    action.description = result['action']['description'];
    action.imageId = result['action-image-id'];
    player.action = action;
  }

  return player;
}
