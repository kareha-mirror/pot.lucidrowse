const List<String> seasonName = ['春', '初夏', '夏', '秋', '初冬', '冬'];

String formatSeason(int day) {
  final int dayOfYear = day % (24 * 6);
  final int month = dayOfYear ~/ 24;

  return seasonName[month];
}
