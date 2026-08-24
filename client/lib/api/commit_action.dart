import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/constants.dart';

Future<Map<String, dynamic>> apiCommitAction(String playerId) async {
  final uri = Uri.parse('$apiBaseUrl/commit-action');

  final response = await http.post(
    uri,
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({'player-id': playerId}),
  );

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  final json = jsonDecode(response.body);
  return json;
}
