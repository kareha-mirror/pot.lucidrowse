class Flavor {
  String input = '';

  String name = '';
  String race = '';
  String job = '';

  String description = '';

  String areaCode = '';
  String areaName = '';

  String? imageId;

  int day = 0;

  bool get hasDescription => description != '';

  Flavor clone() {
    final flavor = Flavor();

    flavor.input = input;
    flavor.name = name;
    flavor.race = race;
    flavor.job = job;
    flavor.description = description;
    flavor.areaCode = areaCode;
    flavor.areaName = areaName;
    flavor.imageId = imageId;
    flavor.day = day;

    return flavor;
  }

  String formatText() {
    final buf = StringBuffer();

    buf.write('名前: $name\n');
    buf.write('種族: $race\n');
    buf.write('職業: $job\n');
    buf.write('\n');
    buf.write('特徴の説明: $description\n');
    buf.write('\n');
    //buf.write('住んでいる地域のコード: $areaCode\n');
    buf.write('住んでいる地域: $areaName\n');

    return buf.toString();
  }
}

class PlayerAction {
  String input = '';

  String description = '';

  String? imageId;

  int day = 0;

  bool get hasDescription => description != '';
}

class Player {
  bool inhabitant = false;
  String id = '';
  int day = 0;

  Flavor flavor = Flavor();

  PlayerAction action = PlayerAction();
  bool committed = false;
}
