import 'package:client/api/post.dart';

Future<void> apiCommitAction() async {
  await apiPost('players/actions/commit', {});
}
