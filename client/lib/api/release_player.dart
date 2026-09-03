import 'package:client/api/post.dart';

Future<void> apiReleasePlayer() async {
  await apiPost('players/release', {});
}
