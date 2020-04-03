package dataobj

var CheckReg = []string{`<div\sid=.*?>.*?</div>\s{0,2}<script>`,
`<div\s*style=.{0,1}position.*?(?:top|left):\s*-[\d]{3,4}px.*?>.*?</div>`,
`<MARQUEE\s.*?scrollAmount=.?[\d]{4,5}.?.*?(?:width|height)=.?[0-5].?.*?>.*?</marquee>`,
`<div\s*style=.?text-indent:\s*-[\d]{2,5}px.?>.*?</div>`,
`<div\s*style=.*?position:\s*absolute\s*;\s*(?:top|left)\s*:\s*expression\(.*?\).*?>.*?</div>`,
`<MARQUEE[^>]*?width=["\']?[0-9]?\s+height=["\']?[0-9]["\']?[^>]*?>([\s\S]*?)</MARQUEE>`,
`<marquee\s+height=[0-9]\s+width=[0-9][^>]*?>([\S\s]*?)</marquee>`,
`<div\s+style\s*=\s*["\']*\s*overflow\s*:\s*hidden\s*;\s*height\s*:\d\d?px\s*;\s*width\s*:\s*\d\d?.*?>([\S\s]*?)</div>`,}


