import 'dart:convert';

import 'package:http/http.dart' as http;

Future<Map<String, dynamic>> apiCreate(String input) async {
  try {
    final uri = Uri.parse('http://localhost:8080/api/player/create');

    final response = await http.post(
      uri,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'input': input}),
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
