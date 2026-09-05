import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';
import 'package:client/models/user.dart';
import 'package:client/state/api_state.dart';

class Api {
  Future<Map<String, dynamic>> get(String path) async {
    late http.Response response;

    try {
      response = await http.get(Uri.parse('$apiBase/$path'));
    } catch (e) {
      apiState.value = ApiState.disconnected;
      rethrow;
    }

    if (response.statusCode == 502) {
      apiState.value = ApiState.disconnected;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    if (response.statusCode == 401) {
      apiState.value = ApiState.unauthorized;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    if (response.statusCode != 200) {
      apiState.value = ApiState.serverError;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    apiState.value = ApiState.connected;
    return jsonDecode(response.body);
  }

  Future<Map<String, dynamic>> post(
    String path,
    Map<String, dynamic> req,
  ) async {
    late http.Response response;

    try {
      response = await http.post(
        Uri.parse('$apiBase/$path'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(req),
      );
    } catch (e) {
      apiState.value = ApiState.disconnected;
      rethrow;
    }

    if (response.statusCode == 502) {
      apiState.value = ApiState.disconnected;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    if (response.statusCode == 401) {
      apiState.value = ApiState.unauthorized;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    if (response.statusCode != 200) {
      apiState.value = ApiState.serverError;
      throw Exception('HTTP error: ${response.statusCode}');
    }

    apiState.value = ApiState.connected;
    return jsonDecode(response.body);
  }

  Future<String> hello() async {
    final result = await get('hello');
    return result['message'];
  }

  Future<Map<String, dynamic>> loadState() async {
    return get('state');
  }

  Future<String> nextDay() async {
    final result = await post('next-day', {});
    return result['error'];
  }

  Future<User> loadUser() async {
    final result = await get('users');
    return User.fromJson(result);
  }

  Future<void> ensureSession() async {
    await post('users/ensure-session', {});
  }

  Future<Player?> loadPlayer() async {
    final result = await get('players');

    if (result['id'] == null) {
      return null;
    }

    final player = Player(id: result['id'], day: result['day']);
    player.points = result['points'];
    player.pointsToUpdate = result['points-to-update'];

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

  Future<Flavor> newFlavor(String input) async {
    final result = await post('players/flavor', {'input': input});
    return Flavor.fromJson(result);
  }

  Future<String> imageFlavor() async {
    final result = await post('players/flavor/image', {});
    return result['image-id'];
  }

  Future<Flavor> updateFlavor(String input) async {
    final result = await post('players/flavor/update', {'input': input});
    return Flavor.fromJson(result);
  }

  Future<void> commitFlavor() async {
    await post('players/flavor/commit', {});
  }

  Future<PlayerAction> newAction(String input) async {
    final result = await post('players/actions', {'input': input});
    return PlayerAction.fromJson(result);
  }

  Future<String> imageAction() async {
    final result = await post('players/actions/image', {});
    return result['image-id'];
  }

  Future<void> commitAction() async {
    await post('players/actions/commit', {});
  }

  Future<Map<String, dynamic>> listPlayers(String regionCode) async {
    return get('regions/$regionCode/players');
  }

  Future<Map<String, dynamic>> listActions(String playerId) async {
    return get('players/$playerId/actions');
  }

  Future<String?> regionState(String regionCode) async {
    final result = await get('regions/$regionCode/state');
    return result['state'];
  }

  Future<void> releasePlayer() async {
    await post('players/release', {});
  }

  Future<void> overridePlayer(String playerId) async {
    await post('players/$playerId/override', {});
  }
}

final api = Api();
