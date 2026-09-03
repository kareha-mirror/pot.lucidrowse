import 'package:client/api/get.dart';

Future<String> apiHello() async {
  final result = await apiGet('hello');
  return result['message'];
}
