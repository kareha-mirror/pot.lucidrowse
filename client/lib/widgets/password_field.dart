import 'package:flutter/material.dart';

class PasswordField extends StatefulWidget {
  const PasswordField({
    super.key,
    this.controller,
    this.hintText,
    this.onChanged,
  });

  final TextEditingController? controller;
  final String? hintText;
  final ValueChanged<String>? onChanged;

  @override
  State<PasswordField> createState() => _PasswordFieldState();
}

class _PasswordFieldState extends State<PasswordField> {
  bool _obscure = true;

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: widget.controller,
      obscureText: _obscure,
      autocorrect: false,
      enableSuggestions: false,
      keyboardType: TextInputType.visiblePassword,
      decoration: InputDecoration(
        hintText: widget.hintText,
        suffixIcon: IconButton(
          icon: Icon(_obscure ? Icons.visibility_off : Icons.visibility),
          tooltip: _obscure ? '秘密を表示' : '秘密を隠す',
          onPressed: () {
            setState(() {
              _obscure = !_obscure;
            });
          },
        ),
      ),
      onChanged: widget.onChanged,
    );
  }
}
