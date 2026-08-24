import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/constants.dart';

Future<Map<String, dynamic>> apiListActions(String playerId) async {
  final uri = Uri.parse('$apiBaseUrl/list-actions/$playerId');

  final response = await http.get(uri);

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
