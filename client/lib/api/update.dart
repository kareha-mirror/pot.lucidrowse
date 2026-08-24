import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/constants.dart';

Future<Map<String, dynamic>> apiUpdate(String id, String input) async {
  try {
    final uri = Uri.parse('$apiBaseUrl/player/update');

    final response = await http.post(
      uri,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'id': id, 'input': input}),
    );

    if (response.statusCode != 200) {
      print('HTTP error: ${response.statusCode}');
      return {};
    }

    final json = jsonDecode(response.body);
    return json;
  } catch (e) {
    print('Connection error: $e');
    return {};
  }
}
