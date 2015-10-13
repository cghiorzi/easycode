package easycode;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class EasycodeController {

    @RequestMapping("/")
    public String greeting() {
        return "index";
    }

}