package easycode;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.ui.Model;

@Controller
public class EasycodeController {

	public static String getName() {
		return "Philadelphia, PA";
	}
	
    @RequestMapping("/")
    public String greeting(Model model) {
		model.addAttribute("location",getName());
        return "index";
    }

}