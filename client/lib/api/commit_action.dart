import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';

Future<Map<String, dynamic>> apiCommitAction(String playerId) async {
  final response = await http.post(
    Uri.parse('$apiBase/players/$playerId/actions/commit'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({}),
  );

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
