class Area {
  final String name;

  const Area({required this.name});
}

class Region {
  final String code;
  final String name;
  final List<Area> areas;

  const Region({required this.code, required this.name, required this.areas});

  String get imagePath => 'assets/images/regions/$code.webp';
  String get thumbnailPath => 'assets/images/regions/$code-thumb.webp';
}

const List<Region> regions = [
  Region(
    code: 'capital',
    name: '王都地方',
    areas: [
      Area(name: '王都'),
      Area(name: '近郊'),
    ],
  ),
  Region(
    code: 'west',
    name: '西方地方',
    areas: [
      Area(name: '街道'),
      Area(name: '丘陵'),
    ],
  ),
  Region(
    code: 'north',
    name: '北方地方',
    areas: [
      Area(name: '山岳'),
      Area(name: '高原'),
    ],
  ),
  Region(
    code: 'east',
    name: '東方地方',
    areas: [
      Area(name: '森林'),
      Area(name: '湖沼'),
    ],
  ),
  Region(
    code: 'south',
    name: '南方地方',
    areas: [
      Area(name: '平野'),
      Area(name: '農村'),
    ],
  ),
  Region(
    code: 'coast',
    name: '海岸地方',
    areas: [
      Area(name: '港湾'),
      Area(name: '海岸'),
    ],
  ),
  Region(
    code: 'islands',
    name: '島嶼地方',
    areas: [
      Area(name: '大島'),
      Area(name: '周辺諸島'),
    ],
  ),
];
