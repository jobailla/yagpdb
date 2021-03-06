package common

import (
	"regexp"
)

type InviteSource struct {
	Name  string
	Regex *regexp.Regexp
}

var DiscordInviteSource = &InviteSource{
	Name:  "Discord",
	Regex: regexp.MustCompile(`(discord\.gg|discordapp\.com\/invite)(?:\/#)?\/([a-zA-Z0-9-]+)`),
}

var ThirdpartyDiscordSites = []*InviteSource{
	&InviteSource{Name: "discord.me", Regex: regexp.MustCompile(`discord\.me\/.+`)},
	&InviteSource{Name: "invite.gg", Regex: regexp.MustCompile(`invite\.gg\/.+`)},
	&InviteSource{Name: "discord.io", Regex: regexp.MustCompile(`discord\.io\/.+`)},
	&InviteSource{Name: "disboard.org", Regex: regexp.MustCompile(`disboard\.org\/server\/join\/.+`)},
	&InviteSource{Name: "discordy.com", Regex: regexp.MustCompile(`discordy\.com\/server\.php`)},

	// regexp.MustCompile(`disco\.gg\/.+`), Youc can't actually link to specific servers here can you, so not needed for now?
}

var AllInviteSources = append([]*InviteSource{DiscordInviteSource}, ThirdpartyDiscordSites...)

func ReplaceServerInvites(msg string, guildID int64, replacement string) string {

	for _, s := range AllInviteSources {
		msg = s.Regex.ReplaceAllString(msg, replacement)
	}

	return msg
}

func ContainsInvite(s string, checkDiscordSource, checkThirdPartySources bool) *InviteSource {
	for _, source := range AllInviteSources {
		if source == DiscordInviteSource && !checkDiscordSource {
			continue
		} else if source != DiscordInviteSource && !checkThirdPartySources {
			continue
		}

		if source.Regex.MatchString(s) {
			return source
		}
	}

	return nil
}
