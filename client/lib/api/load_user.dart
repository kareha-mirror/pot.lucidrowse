import 'package:client/api/get.dart';
import 'package:client/models/user.dart';

Future<User> apiLoadUser() async {
  final result = await apiGet('users');
  return User.fromJson(result);
}
