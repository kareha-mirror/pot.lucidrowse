import 'dart:convert';

import 'package:http/http.dart' as http;

Future<Map<String, dynamic>> apiHello() async {
  final uri = Uri.parse('http://localhost:8080/api/hello');

  final response = await http.get(uri);

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
