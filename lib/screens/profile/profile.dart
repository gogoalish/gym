import "package:flutter/material.dart";

import "../../widgets/MeCard.dart";

class Profile extends StatelessWidget {
  const Profile({
    super.key,
  });
  @override
  Widget build(BuildContext context) {
    return Container(
        padding: const EdgeInsets.all(16),
        child: ListView(
          children: [
            Container(
              margin: const EdgeInsets.only(bottom: 24),
              child: const Text(
                "My profile",
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 32,
                ),
              ),
            ),
            Row(
              children: [
                ClipRRect(
                  borderRadius: BorderRadius.circular(10.0),
                  child: Container(
                      decoration: const BoxDecoration(
                        color: Color(0xffFF0E58),
                      ),
                      height: 60,
                      width: 60,
                      child: Image.asset("assets/images/feitan.jpg")),
                ),
                Container(
                  margin: const EdgeInsets.only(left: 16),
                  height: 60,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    children: const [
                      Text(
                        "Sanzhar",
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 20,
                        ),
                      ),
                      Text(
                        "sanzhar.sseiful@gmail.com",
                        style: TextStyle(
                          fontSize: 16,
                        ),
                      )
                    ],
                  ),
                )
              ],
            ),
            const Padding(padding: EdgeInsets.all(24)),
            const MeCard(
              title: "My courses",
              hint: "already have 3 courses",
            ),
            // const MeCard(
            //   title: "Workouts",
            //   hint: "3 workouts",
            // ),
            const MeCard(
              title: "Promocodes",
              hint: "Claim your promocodes",
            ),
            const MeCard(
              title: "Ask and answer",
              hint: "Ask your questions..",
            ),
            const MeCard(
              title: "Payment methods",
              hint: "Visa *4565",
            ),
            const MeCard(
              title: "Settings",
              hint: "Change your settings",
            ),
          ],
        ));
  }
}
