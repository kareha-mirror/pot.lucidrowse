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
  .new(
    code: 'capital',
    name: '王都地方',
    areas: [
      .new(name: '王都'),
      .new(name: '近郊'),
    ],
  ),
  .new(
    code: 'west',
    name: '西方地方',
    areas: [
      .new(name: '街道'),
      .new(name: '丘陵'),
    ],
  ),
  .new(
    code: 'north',
    name: '北方地方',
    areas: [
      .new(name: '山岳'),
      .new(name: '高原'),
    ],
  ),
  .new(
    code: 'east',
    name: '東方地方',
    areas: [
      .new(name: '森林'),
      .new(name: '湖沼'),
    ],
  ),
  .new(
    code: 'south',
    name: '南方地方',
    areas: [
      .new(name: '平野'),
      .new(name: '農村'),
    ],
  ),
  .new(
    code: 'coast',
    name: '海岸地方',
    areas: [
      .new(name: '港湾'),
      .new(name: '海岸'),
    ],
  ),
  .new(
    code: 'islands',
    name: '島嶼地方',
    areas: [
      .new(name: '大島'),
      .new(name: '周辺諸島'),
    ],
  ),
];
