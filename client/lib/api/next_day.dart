import 'package:client/api/post.dart';

Future<String> apiNextDay() async {
  final result = await apiPost('next-day', {});
  return result['error'];
}
