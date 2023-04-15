import "package:flutter/material.dart";

class MusicCard extends StatelessWidget {
  final String name;
  final String image;
  const MusicCard({super.key, required this.name, required this.image});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(vertical: 8),
      child: Row(children: [
        Image.network(
          "$image",
          width: 60,
          height: 60,
        ),
        Padding(padding: EdgeInsets.only(right: 8)),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                overflow: TextOverflow.ellipsis,
                "$name",
                style: TextStyle(
                  fontSize: 22,
                  fontWeight: FontWeight.bold,
                ),
              ),
              Row(
                children: [
                  Container(
                    padding: EdgeInsets.all(4),
                    margin: EdgeInsets.only(right: 4),
                    decoration: BoxDecoration(
                      color: Colors.grey,
                      borderRadius: BorderRadius.circular(4),
                    ),
                    child: Text(
                      "Lyrics",
                      style: TextStyle(
                        color: Colors.white,
                      ),
                    ),
                  ),
                  Text(
                    "Safe and Sound",
                    style: TextStyle(),
                  ),
                ],
              ),
            ],
          ),
        ),
        Icon(Icons.more_vert),
      ]),
    );
  }
}
