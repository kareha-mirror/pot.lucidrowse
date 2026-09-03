import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';

Future<int> apiLoadDay() async {
  final response = await http.get(Uri.parse('$apiBase/day'));

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  final result = jsonDecode(response.body);
  return result['day'];
}
