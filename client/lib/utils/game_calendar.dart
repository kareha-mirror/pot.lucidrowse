const List<String> monthName = ['芽吹き', '若葉', '陽盛り', '実り', '霜降り', '雪籠り'];

String formatGameDate(int day) {
  final int year = day ~/ (24 * 6);
  final int dayOfYear = day % (24 * 6);
  final int month = dayOfYear ~/ 24;
  final int dayOfMonth = dayOfYear % 24;

  return '王都暦${year + 568}年 ${monthName[month]}の月 ${dayOfMonth + 1}日';
}
