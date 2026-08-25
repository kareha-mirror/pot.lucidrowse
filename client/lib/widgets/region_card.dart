import 'package:flutter/material.dart';

import 'package:client/models/region.dart';

class RegionCard extends StatelessWidget {
  const RegionCard({super.key, required this.region, this.onTap});

  final Region region;
  final GestureTapCallback? onTap;

  @override
  Widget build(BuildContext context) {
    final List<String> areas = [];
    for (final area in region.areas) {
      areas.add(area.name);
    }

    return Card(
      clipBehavior: Clip.antiAlias,
      child: InkWell(
        onTap: onTap,
        child: Row(
          children: [
            Image.asset(
              region.image,
              width: MediaQuery.sizeOf(context).width < 600 ? 150 : 300,
              height: MediaQuery.sizeOf(context).width < 600 ? 75 : 150,
              fit: BoxFit.cover,
            ),

            const SizedBox(width: 16),

            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    region.name,
                    style: Theme.of(context).textTheme.titleMedium,
                  ),

                  const SizedBox(height: 4),

                  Text(areas.join(' / ')),
                ],
              ),
            ),

            const SizedBox(width: 12),
          ],
        ),
      ),
    );
  }
}
