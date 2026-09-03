import 'package:client/api/post.dart';

Future<void> apiCommitFlavor() async {
  await apiPost('players/flavor/commit', {});
}
