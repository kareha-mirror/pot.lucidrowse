import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/constants.dart';

Future<Map<String, dynamic>> apiNewAction(String playerId, String input) async {
  final uri = Uri.parse('$apiBaseUrl/new-action');

  final response = await http.post(
    uri,
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({'player-id': playerId, 'input': input}),
  );

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  final json = jsonDecode(response.body);
  return json;
}
