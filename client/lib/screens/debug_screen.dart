import 'dart:typed_data';
import 'package:flutter/material.dart';

import 'package:client/api/health.dart';
import 'package:client/api/image.dart';
import 'package:client/api/interpret.dart';
import 'package:client/models/player.dart';
import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/widgets/translucent_panel.dart';

class DebugScreen extends StatefulWidget {
  final AppState state;

  const DebugScreen({super.key, required this.state});

  @override
  State<DebugScreen> createState() => _DebugScreenState();
}

class _DebugScreenState extends State<DebugScreen> {
  late TextEditingController _controller;
  late TextEditingController _filtered;
  bool _textLoading = false;
  bool _imageLoading = false;
  Flavor _flavor = Flavor();
  Uint8List? imageData;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
    _filtered = TextEditingController();
  }

  @override
  void dispose() {
    _controller.dispose();
    _filtered.dispose();
    super.dispose();
  }

  void _nextDay() {
    setState(() {
      for (final inhabitant in widget.state.inhabitants) {
        if (inhabitant.committed) {
          inhabitant.actions.add(inhabitant.action);
        }
        inhabitant.action = PlayerAction();
      }

      widget.state.day++;
    });
  }

  void _clearMyself() {
    setState(() => widget.state.player = Player());
  }

  void _clearInhabitants() {
    setState(() {
      widget.state.player = Player();
      widget.state.inhabitants = [];
    });
  }

  Future<void> _fetchImage(String id) async {
    setState(() => _imageLoading = true);
    try {
      final img = await fetchImage(id);

      setState(() {
        imageData = img;
      });
    } catch (e) {
      // TODO
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _interpretCharacter() async {
    setState(() => _textLoading = true);
    try {
      final filtered = await interpretCharacter(_controller.text);
      final f = Flavor();
      f.name = filtered['flavor']['name'];
      f.race = filtered['flavor']['race'];
      f.job = filtered['flavor']['job'];
      f.filtered = filtered['flavor']['text'];
      setState(() {
        _flavor = f;
        _filtered.text = _flavor.formatText();
      });

      _fetchImage(filtered['id']);
    } catch (e) {
      // TODO
    } finally {
      setState(() => _textLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          SizedBox(
            width: double.infinity,
            height: double.infinity,
            child: const Image(
              image: AssetImage('assets/images/debug.webp'),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  TranslucentPanel(
                    child: Column(
                      children: [
                        Text(formatDate(widget.state.day)),
                        Text('夢路開通 ${widget.state.day + 1} 日目'),
                      ],
                    ),
                  ),
                  const SizedBox(height: 12),
                  ElevatedButton(
                    onPressed: () => _nextDay(),
                    child: const Text('明日まで寝る'),
                  ),

                  const SizedBox(height: 48),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: widget.state.player.isForeigner
                          ? null
                          : () => _clearMyself(),
                      child: const Text('自分を手放す'),
                    ),
                  ),

                  const SizedBox(height: 48),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: widget.state.inhabitants.isEmpty
                          ? null
                          : () => _clearInhabitants(),
                      child: const Text('住人たちを追い出す'),
                    ),
                  ),

                  const SizedBox(height: 48),

                  ElevatedButton(
                    onPressed: () async {
                      await checkHealth();
                    },
                    child: const Text('Check Health'),
                  ),

                  const SizedBox(height: 12),

                  TranslucentPanel(
                    child: TextField(
                      controller: _controller,
                      decoration: InputDecoration(
                        hintText: '(どんな存在になりたいか、ここに念じよう。)',
                      ),
                      maxLines: null,
                      maxLength: 140,
                      readOnly: _textLoading || _imageLoading,
                    ),
                  ),

                  const SizedBox(height: 12),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: (_textLoading || _imageLoading)
                          ? null
                          : _interpretCharacter,
                      child: const Text('Interpret Character'),
                    ),
                  ),

                  const SizedBox(height: 12),

                  if (_textLoading)
                    const CircularProgressIndicator()
                  else
                    TranslucentPanel(
                      child: TextField(
                        controller: _filtered,
                        decoration: InputDecoration(
                          hintText: '(夢に映され、こうなりました。)',
                        ),
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                  const SizedBox(height: 12),

                  if (_imageLoading)
                    const CircularProgressIndicator()
                  else if (imageData != null)
                    ConstrainedBox(
                      constraints: const BoxConstraints(maxWidth: 300),
                      child: ClipRRect(
                        borderRadius: BorderRadius.circular(12),
                        child: Image.memory(imageData!),
                      ),
                    ),

                  const SizedBox(height: 48),

                  helpButton(context, HelpPage.debug),

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: const Text('ふたを閉じる'),
                  ),

                  const SizedBox(height: 96),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
