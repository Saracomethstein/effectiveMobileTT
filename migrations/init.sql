CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE songs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    release_date DATE NOT NULL,
    lyrics TEXT NOT NULL,
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO songs (group_name, song_name, release_date, lyrics, link)
VALUES
    ('Kendrick Lamar', 'tv off', '2024-11-22',
     'All I ever wanted was a black Grand National
      Fuck being rational, give ''em what they ask for
      It''s not enough (ayy)
      Few solid niggas left, but it''s not enough
      Few bitches that''ll really step, but it''s not enough
      Say you bigger than myself, but it''s not enough (huh)
      I get on they ass, yeah, somebody gotta do it
      I''ll make them niggas mad, yeah, somebody gotta do it
      I''ll take they G-pass, shit, watch a nigga do it
      Huh, we survived outside, all from the music, nigga, what?
      They like, "What he on?"
      It''s the Alpha and Omega, bitch, welcome home
      This is not a song
      This a revelation, how to get a nigga gone
      You need you a man, baby, I don''t understand, baby
      Pay your bill and make you feel protected like I can, baby
      Teach you someth''in if you need correction, that''s the plan, baby
      Don''t put your life in these weird niggas'' hands, baby (whoa)
      It''s not enough (ayy)
      Few solid niggas left, but it''s not enough
      Few bitches that''ll really step, but it''s not enough
      Say you bigger than myself, but it''s not enough (huh)
      I get on they ass, yeah, somebody gotta do it
      I''ll make them niggas mad, yeah, somebody gotta do it
      I''ll take they G-pass, shit, watch a nigga do it
      Huh, we survived outside, all from the music, nigga, what?
      Hey, turn his TV off
      Ain''t with my type activities? Then don''t you get involved
      Hey, what, huh, how many should I send? Send ''em all
      Take a risk or take a trip, you know I''m trippin'' for my dawg
      Who you with? Couple sergeants and lieutenants for the get-back
      This revolution been televised, I fell through with the knick-knacks
      Hey, young nigga, get your chili up, yeah, I meant that
      Hey, black out if they act out, yeah, I did that
      Hey, what''s up, though?
      I hate a bitch that''s hatin'' on a bitch and they both hoes
      I hate a nigga hatin'' on them niggas and they both broke
      If you ain''t coming for no chili, what you come for?
      Nigga feel like he entitled ''cause he knew me since a kid
      Bitch, I cut my granny off if she don''t see it how I see it, hm
      Got a big mouth but he lack big ideas
      Send him to the moon, that''s just how I feel, yellin''
      It''s not enough (ayy)
      Few solid niggas left, but it''s not enough
      Few bitches that''ll really step, but it''s not enough
      Say you bigger than myself, but it''s not enough
      Huh, huh, huh
      Hey
      Hey (Mustard on the beat, ho)
      Mustard
      Niggas actin'' bad, but somebody gotta do it
      Got my foot up on the gas, but somebody gotta do it, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off
      Ain''t no other king in this rap thing, like siblings
      Nothing but my children, one shot, they disappearin''
      I''m in a city with a flag
      Be gettin'' thrown like it was pass interference
      Padlock around the building
      Crash, pullin'' up in unmarked truck just to play freeze tag
      With a bone to pick like it was sea bass
      So when I made it out, I made about 50K from a show
      Tryna show niggas the ropes before they hung from a rope
      I''m prophetic, they only talk about it, I get it
      Only good for saving face, seen the cosmetics
      How many heads I gotta take to level my aesthetics?
      Hurry up and get your muscle up, we out the plyometric
      Nigga ran up out of lux soon as I up the highest metric
      The city just made it sweet, you could die, bet it
      They mouth get full of deceit, let these cowards tell it
      Walk in New Orleans with the etiquette of L.A., yellin''
      Mustard (oh, man)
      Niggas actin'' bad, but somebody gotta do it
      Got my foot up on the gas, but somebody gotta do it, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off, huh
      Turn his TV off, turn his TV off
      Shit gets crazy, scary, spooky, hilarious
      Crazy, scary, spooky, hilarious
      Shit gets crazy, scary, spooky, hilarious
      Crazy, scary, spooky, hilarious
      Shit gets crazy, scary, spooky, hilarious
      Crazy, scary, spooky, hilarious
      Shit gets crazy, scary, spooky, hilarious
      Crazy, scary, spooky, hilarious',
     'https://www.youtube.com/watch?v=U8F5G5wR1mk&ab_channel=KendrickLamar');
