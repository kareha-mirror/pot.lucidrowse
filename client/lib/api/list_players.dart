import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/constants.dart';

Future<Map<String, dynamic>> apiListPlayers() async {
  final uri = Uri.parse('$apiBaseUrl/list-players');

  final response = await http.get(uri);

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
