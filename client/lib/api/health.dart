import 'dart:convert';

import 'package:http/http.dart' as http;

Future<void> checkHealth() async {
  try {
    final uri = Uri.parse('http://localhost:8080/api/health');

    final response = await http.get(uri);

    if (response.statusCode != 200) {
      print('HTTP error: ${response.statusCode}');
      return;
    }

    final json = jsonDecode(response.body);
    print(json);
  } catch (e) {
    print('Connection error: $e');
  }
}
