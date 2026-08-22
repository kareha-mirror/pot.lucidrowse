import 'dart:convert';

import 'package:http/http.dart' as http;

Future<String> interpretCharacter(String text) async {
  try {
    final uri = Uri.parse('http://localhost:8080/api/character/interpret');

    final response = await http.post(
      uri,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'text': text}),
    );

    if (response.statusCode != 200) {
      print('HTTP error: ${response.statusCode}');
      return '';
    }

    final json = jsonDecode(response.body);
    return json['text'];
  } catch (e) {
    print('Connection error: $e');
    return '';
  }
}
