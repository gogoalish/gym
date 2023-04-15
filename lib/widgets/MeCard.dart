import 'package:flutter/material.dart';

class MeCard extends StatelessWidget {
  final String title;
  final String hint;

  const MeCard({
    super.key,
    required this.title,
    required this.hint,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(bottom: 16, top: 16),
      decoration: const BoxDecoration(
        border: Border(
          bottom: BorderSide(color: Color.fromARGB(255, 238, 235, 235)),
        ),
      ),
      child: Row(
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                padding: const EdgeInsets.only(bottom: 8),
                child: Text(
                  "$title",
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 18,
                  ),
                ),
              ),
              Text(
                "$hint",
                style: const TextStyle(fontSize: 14, color: Colors.grey),
              )
            ],
          ),
          const Spacer(),
          const Icon(
            Icons.arrow_right,
            size: 36,
            color: Colors.grey,
          )
        ],
      ),
    );
  }
}
