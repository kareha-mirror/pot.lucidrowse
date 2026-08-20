import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';

class CreateScreen extends StatefulWidget {
  final AppState state;

  const CreateScreen({super.key, required this.state});

  @override
  State<CreateScreen> createState() => _CreateScreenState();
}

class _CreateScreenState extends State<CreateScreen> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
    _controller.text = widget.state.player.flavor.raw;
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _submitRaw(String value) {
    setState(() {
      final flavor = widget.state.player.flavor;

      flavor.raw = value;
      flavor.filtered = value;

      flavor.name = 'リタ';
      flavor.race = '人間';
      flavor.job = '見習い魔法使い';

      flavor.imageUrl = 'assets/images/figure.webp';

      flavor.day = widget.state.day;
    });
  }

  void _commit() {
    setState(() {
      final player = widget.state.player;
      player.flavors.add(player.flavor);

      if (player.inhabit.isForeigner) {
        player.inhabit = Inhabit.inhabitant;
        player.settled = widget.state.day;
        widget.state.inhabitants.add(player);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('まどろみの水晶球'),
      ),

      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Center(
          child: Column(
            children: [
              if (widget.state.player.inhabit.isForeigner) ...[
                const Text('水晶球が心に呼びかける。'),
                const Text('夢の世界へいざないましょう。'),
                const SizedBox(height: 24),
                const Text('あなたはこの世界に自身の分身を作り出し、夢と関わっていくことになります。'),
                const Text('この世界で夢見るあなたは、どんな人ですか、どんな生きものですか。'),
              ] else ...[
                const Text('水晶球が心に呼びかける。'),
                const Text('何か変化はありましたか。'),
                const SizedBox(height: 24),
                const Text('あなたの分身に何か変化があるのなら、ここで新たに念じてください。'),
              ],

              const SizedBox(height: 24),

              Padding(
                padding: EdgeInsets.symmetric(horizontal: 48),
                child: TextField(
                  controller: _controller,
                  decoration: const InputDecoration(hintText: '森の薬草取りの少女。'),
                  onSubmitted: (String value) => _submitRaw(value),
                  onChanged: (String value) => setState(() {}),
                  maxLines: null,
                  maxLength: 140,
                ),
              ),

              const SizedBox(height: 24),

              ElevatedButton(
                onPressed: _controller.text.isEmpty
                    ? null
                    : () => _submitRaw(_controller.text),
                child: const Text('念じる'),
              ),

              if (widget.state.player.flavor.filtered != '') ...[
                const SizedBox(height: 48),

                const Text('あなたの言葉は夢に映され、こうなりました。'),

                const SizedBox(height: 24),

                Text(widget.state.player.flavor.filtered),

                const SizedBox(height: 24),

                Text('名前 : ${widget.state.player.flavor.name}'),
                Text('種族 : ${widget.state.player.flavor.race}'),
                Text('職業 : ${widget.state.player.flavor.job}'),

                const SizedBox(height: 48),

                Center(
                  child: ConstrainedBox(
                    constraints: const BoxConstraints(maxWidth: 600),
                    child: ClipRRect(
                      borderRadius: BorderRadius.circular(24),
                      child: Image(
                        image: AssetImage(widget.state.player.flavor.imageUrl),
                      ),
                    ),
                  ),
                ),

                const SizedBox(height: 48),

                ElevatedButton(
                  onPressed: () {
                    _commit();
                    Navigator.pop(context);
                  },
                  child: const Text('これでよし'),
                ),
              ],

              const SizedBox(height: 96),

              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('目をそらす'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
