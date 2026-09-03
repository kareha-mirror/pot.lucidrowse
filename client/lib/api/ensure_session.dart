import 'package:client/api/post.dart';

Future<void> apiEnsureSession() async {
  await apiPost('users/ensure-session', {});
}
